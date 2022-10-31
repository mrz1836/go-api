// Package database provides a layer for interacting with read/write databases
package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/mrz1836/go-logger"
	// used for the mysql driver
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

// Database constants
const (
	MySQLDriver      = "mysql"
	PostgreSQLDriver = "postgresql"
)

// Global database instances
var (
	config        Configuration
	ReadDatabase  *APIDatabase
	WriteDatabase *APIDatabase
)

// APIDatabase Extends sql.DB
type APIDatabase struct {
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
	DatabaseRead  ConnectionConfig `json:"database_read" mapstructure:"database_read"`   // Read database connection
	DatabaseWrite ConnectionConfig `json:"database_write" mapstructure:"database_write"` // Write database connection
}

// ConnectionConfig is a configuration for a SQL connection
type ConnectionConfig struct {
	Driver             string `json:"driver" mapstructure:"driver"`                             // mysql or postgresql
	Host               string `json:"host" mapstructure:"host"`                                 // localhost
	MaxConnectionTime  int    `json:"max_connection_time" mapstructure:"max_connection_time"`   // 60
	MaxIdleConnections int    `json:"max_idle_connections" mapstructure:"max_idle_connections"` // 5
	MaxOpenConnections int    `json:"max_open_connections" mapstructure:"max_open_connections"` // 5
	Name               string `json:"name" mapstructure:"name"`                                 // database-name
	Password           string `json:"password" mapstructure:"password"`                         // user-password
	Port               string `json:"port" mapstructure:"port"`                                 // 3306
	User               string `json:"user" mapstructure:"user"`                                 // username
}

// SetConfiguration sets the configuration
func SetConfiguration(conf Configuration) {
	config = conf
}

// NewAPIDatabase creates a new database connection
func NewAPIDatabase(read, write *sql.DB) *APIDatabase {
	databaseQueue := &APIDatabase{read, write, nil, nil, nil, nil, nil}
	databaseQueue.throttleQueue = make(chan struct{}, 10)
	databaseQueue.statements = make(map[uint32]*sql.Stmt, 30)
	databaseQueue.worker = make(chan func(), 10000)
	databaseQueue.statementMutex = new(sync.RWMutex)
	databaseQueue.waitGroup = new(sync.WaitGroup)
	go databaseQueue.startWorker()
	return databaseQueue
}

// OpenConnection opens the database connection (read / write)
func OpenConnection() (err error) {
	var dB *sql.DB      // I am read only but cheap.  Please use me when possible.
	var dBWrite *sql.DB // I am read-write, but expensive.  Please only use me for writing.

	// Open the main database connection (read)
	err = openDatabaseConnection(
		config.DatabaseRead.Host+":"+config.DatabaseRead.Port,
		config.DatabaseRead.Name,
		config.DatabaseRead.User,
		config.DatabaseRead.Password,
		&dB,
		config.DatabaseRead.Driver,
		config.DatabaseRead.MaxOpenConnections,
		config.DatabaseRead.MaxIdleConnections,
		config.DatabaseRead.MaxConnectionTime,
	)
	if err != nil {
		return
	}

	// Open a write connection if found and different from the read
	if len(config.DatabaseWrite.Host) == 0 || (config.DatabaseWrite.Host+":"+config.DatabaseWrite.Port) == (config.DatabaseRead.Host+":"+config.DatabaseRead.Port) {
		dBWrite = dB
	} else {
		err = openDatabaseConnection(
			config.DatabaseWrite.Host+":"+config.DatabaseWrite.Port,
			config.DatabaseWrite.Name,
			config.DatabaseWrite.User,
			config.DatabaseWrite.Password,
			&dBWrite,
			config.DatabaseWrite.Driver,
			config.DatabaseWrite.MaxOpenConnections,
			config.DatabaseWrite.MaxIdleConnections,
			config.DatabaseWrite.MaxConnectionTime,
		)
		if err != nil {
			return
		}
	}

	// Set the new connections
	ReadDatabase = NewAPIDatabase(dB, dBWrite)
	WriteDatabase = NewAPIDatabase(dBWrite, dBWrite)

	return
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

// NewTx creates a new TX
func NewTx(timeout time.Duration) (tx *sql.Tx, cancelMethod context.CancelFunc, err error) {
	var ctx context.Context
	ctx, cancelMethod = context.WithTimeout(context.Background(), timeout)
	tx, err = WriteDatabase.BeginTx(ctx, nil)
	return
}

// startWorker starts a worker for the Throttled Query
func (d *APIDatabase) startWorker() {
	for handle := range d.worker {
		handle()
		d.waitGroup.Done()
	}
}

// StopWorker will wait for the queue to empty and then shutdown the worker
func (d *APIDatabase) StopWorker() {
	d.waitGroup.Wait()
	close(d.worker)
}

// Enque adds a worker
func (d *APIDatabase) Enque(handle func()) {
	d.waitGroup.Add(1)
	d.worker <- handle
}

// Close both or any connections
func (d *APIDatabase) Close() {
	_ = d.DB.Close() // todo: log these errors if needed
	if d.dBWrite != d.DB {
		_ = d.dBWrite.Close()
	}
}

// GetReadDatabase gets the read database connection these are needed for testing because
// for some reason it can't determine that DeliveryDudesDB extends sql.DB
func (d *APIDatabase) GetReadDatabase() *sql.DB {
	return d.DB
}

// GetWriteDatabase gets the "write database" connection
func (d *APIDatabase) GetWriteDatabase() *sql.DB {
	return d.dBWrite
}

// openDatabaseConnection opens a new database connection
func openDatabaseConnection(databaseAddress, databaseName, databaseUser, databasePassword string,
	database **sql.DB, driver string, maxOpenConnections, maxOpenIdle, maxConnectionLifetime int) (err error) {

	// Start a connection with the database
	var db *sql.DB

	// Switch on the drivers supported
	switch driver {
	case MySQLDriver:
		db, err = sql.Open(driver, databaseUser+":"+databasePassword+"@tcp("+databaseAddress+")/"+databaseName+"?parseTime=true")
	case PostgreSQLDriver:
		db, err = sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s/%s", databaseUser, databasePassword, databaseAddress, databaseName))
	default:
		logger.Data(2, logger.ERROR, fmt.Sprintf("unknown driver specified: %s", driver))
		return
	}

	// Do we have a fatal error opening
	if err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("database open error: %s at %s", driver, databaseAddress), logger.MakeParameter("db_error", err.Error()))
		return
	}

	// Do we have a valid db
	if db == nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("database connection is nil: %s at %s", driver, databaseAddress))
		return
	}

	// Do we have a db connection?
	if err = db.Ping(); err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("database ping error: %s at %s", driver, databaseAddress), logger.MakeParameter("db_error", err.Error()))
		return
	}

	// Set the max open connections to the DB
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxOpenIdle)

	// Set the max time a connection will be considered open
	db.SetConnMaxLifetime(time.Duration(maxConnectionLifetime) * time.Second)

	// Set the global DB
	*database = db

	// Debug statement for connection
	logger.Data(2, logger.DEBUG, "database: connected",
		logger.MakeParameter("address", databaseAddress),
		logger.MakeParameter("driver", driver),
		logger.MakeParameter("max_idle_connections", maxOpenIdle),
		logger.MakeParameter("max_lifetime", maxConnectionLifetime),
		logger.MakeParameter("max_open_connections", maxOpenConnections),
		logger.MakeParameter("name", databaseName),
		logger.MakeParameter("user", databaseUser),
	)

	return
}
