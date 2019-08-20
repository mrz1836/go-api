// Package database provides a layer for interacting with read/write databases
package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/mrz1836/go-logger"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql/driver"
)

// Database constants
const (
	MySQLDriver      = "mysql"
	PostgreSQLDriver = "postgresql"
)

// Global database instances
var (
	config        Configuration
	ReadDatabase  *ApiDatabase
	WriteDatabase *ApiDatabase
)

// ApiDatabase Extends sql.DB
type ApiDatabase struct {
	*sql.DB                // calls are passed through by default to me
	dBWrite        *sql.DB // I am read-write, but expensive.  Please only use me for writing.
	throttleQueue  chan struct{}
	statements     map[uint32]*sql.Stmt
	statementMutex *sync.RWMutex
	worker         chan func()
	waitGroup      *sync.WaitGroup
	//	dB      *sql.DB // same as super struct; I am read only but cheap.  Please use me when possible.
	// throttleCount int32  // uncomment when adding query weight
}

// Configuration is the database configuration
type Configuration struct {
	DatabaseDriver                  string `env:"DATABASE_DRIVER"`                     //Database driver type (mysql, postgresql)
	DatabaseMaxConnLifetime         int    `env:"DATABASE_MAX_CONN_LIFETIME"`          //Maximum seconds for a database connection
	DatabaseReadHost                string `env:"DATABASE_READ_HOST"`                  //Main database host (read) (ip -- 127.0.0.1)
	DatabaseReadMaxIdleConnections  int    `env:"DATABASE_READ_MAX_IDLE_CONNECTIONS"`  //Maximum amount of idle connections available
	DatabaseReadMaxOpenConnections  int    `env:"DATABASE_READ_MAX_OPEN_CONNECTIONS"`  //Maximum amount of open connections in pool
	DatabaseReadName                string `env:"DATABASE_READ_NAME"`                  //Main database name (read)
	DatabaseReadPort                string `env:"DATABASE_READ_PORT"`                  //Main database port (read)
	DatabaseReadUser                string `env:"DATABASE_READ_USER"`                  //Database username (read)
	DatabaseReadPassword            string `env:"DATABASE_READ_PASSWORD"`              //Database user's password (read)
	DatabaseWriteHost               string `env:"DATABASE_WRITE_HOST"`                 //Main database host (write)
	DatabaseWriteMaxIdleConnections int    `env:"DATABASE_WRITE_MAX_IDLE_CONNECTIONS"` //Maximum amount of idle connections available
	DatabaseWriteMaxOpenConnections int    `env:"DATABASE_WRITE_MAX_OPEN_CONNECTIONS"` //Maximum amount of open connections in pool
	DatabaseWriteName               string `env:"DATABASE_WRITE_NAME"`                 //Main database name (write)
	DatabaseWritePort               string `env:"DATABASE_WRITE_PORT"`                 //Main database port (write)
	DatabaseWriteUser               string `env:"DATABASE_WRITE_USER"`                 //Database username (write)
	DatabaseWritePassword           string `env:"DATABASE_WRITE_PASSWORD"`             //Database user's password (write)
}

// SetConfiguration sets the configuration
func SetConfiguration(conf Configuration) {
	config = conf
}

// NewApiDatabase creates a new database connection
func NewApiDatabase(read, write *sql.DB) *ApiDatabase {
	databaseQueue := &ApiDatabase{read, write, nil, nil, nil, nil, nil}
	databaseQueue.throttleQueue = make(chan struct{}, 10)
	databaseQueue.statements = make(map[uint32]*sql.Stmt, 30)
	databaseQueue.worker = make(chan func(), 10000)
	databaseQueue.statementMutex = new(sync.RWMutex)
	databaseQueue.waitGroup = new(sync.WaitGroup)
	go databaseQueue.startWorker()
	return databaseQueue
}

// OpenConnection opens the database connection (read / write)
func OpenConnection() {
	var dB *sql.DB      // I am read only but cheap.  Please use me when possible.
	var dBWrite *sql.DB // I am read-write, but expensive.  Please only use me for writing.

	//Open the main database connection (read)
	openDatabaseConnection(config.DatabaseReadHost+":"+config.DatabaseReadPort, config.DatabaseReadName, config.DatabaseReadUser, config.DatabaseReadPassword, &dB, config.DatabaseReadMaxOpenConnections, config.DatabaseReadMaxIdleConnections)

	//Open a write connection if found and different from the read
	if len(config.DatabaseWriteHost) == 0 || (config.DatabaseWriteHost+":"+config.DatabaseWritePort) == (config.DatabaseReadHost+":"+config.DatabaseReadPort) {
		dBWrite = dB
	} else {
		openDatabaseConnection(config.DatabaseWriteHost+":"+config.DatabaseWritePort, config.DatabaseWriteName, config.DatabaseWriteUser, config.DatabaseWritePassword, &dBWrite, config.DatabaseWriteMaxOpenConnections, config.DatabaseWriteMaxIdleConnections)
	}
	ReadDatabase = NewApiDatabase(dB, dBWrite)
	WriteDatabase = NewApiDatabase(dBWrite, dBWrite)
}

// CloseAllConnections closes the current database connections
func CloseAllConnections() {
	WriteDatabase.waitGroup.Wait()
	WriteDatabase.Close()
	WriteDatabase = nil
	ReadDatabase.waitGroup.Wait()
	ReadDatabase.Close()
	ReadDatabase = nil
}

// startWorker starts a worker for the Throttled Query
func (d *ApiDatabase) startWorker() {
	for handle := range d.worker {
		handle()
		d.waitGroup.Done()
	}
}

// StopWorker will wait for the queue to empty and then shutdown the worker
func (d *ApiDatabase) StopWorker() {
	d.waitGroup.Wait()
	close(d.worker)
}

// Enque adds a worker
func (d *ApiDatabase) Enque(handle func()) {
	d.waitGroup.Add(1)
	d.worker <- handle
}

//Close both or any connections
func (d *ApiDatabase) Close() {
	_ = d.DB.Close() // todo: log these errors if needed
	if d.dBWrite != d.DB {
		_ = d.dBWrite.Close()
	}
}

// GetReadDatabase gets the read database connection these are needed for testing because
// for some reason it can't determine that DeliveryDudesDB extends sql.DB
func (d *ApiDatabase) GetReadDatabase() *sql.DB {
	return d.DB
}

// GetWriteDatabase gets the write database connection
func (d *ApiDatabase) GetWriteDatabase() *sql.DB {
	return d.dBWrite
}

// openDatabaseConnection opens a new database connection
func openDatabaseConnection(databaseAddress, databaseName, databaseUser, databasePassword string, database **sql.DB, maxOpenConnections, maxOpenIdle int) {
	//Start a connection with the database
	var db *sql.DB
	var err error
	switch config.DatabaseDriver {
	case MySQLDriver:
		db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseAddress+")/"+databaseName+"?parseTime=true")
	case PostgreSQLDriver:
		db, err = sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", databaseUser, databasePassword, databaseAddress, databaseName))
		logger.Printf("postgres://%s:%s@%s/%s", databaseUser, databasePassword, databaseAddress, databaseName)
	default:
		logger.Fatalln("unknown driver specified:", config.DatabaseDriver)
	}

	//Do we have a fatal error opening
	if err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("database open error: %s", databaseAddress), logger.MakeParameter("db_error", err.Error()))
		logger.Fatalln("database open: ", err)
	}

	//Do we have a db connection?
	if err = db.Ping(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("database ping error: %s", databaseAddress), logger.MakeParameter("db_error", err.Error()))
		logger.Fatalln("database ping: ", err)
	}

	//Set the max open connections to the DB
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxOpenIdle)

	//Set the max time a connection will be considered open
	db.SetConnMaxLifetime(time.Duration(config.DatabaseMaxConnLifetime) * time.Second)

	//Set the global DB
	*database = db

	//Done!
	logger.Data(2, logger.INFO, "database: connected",
		logger.MakeParameter("address", databaseAddress),
		logger.MakeParameter("name", databaseName),
		logger.MakeParameter("max_open_connections", maxOpenConnections),
		logger.MakeParameter("max_idle_connections", maxOpenIdle),
		logger.MakeParameter("max_lifetime", config.DatabaseMaxConnLifetime),
	)
}
