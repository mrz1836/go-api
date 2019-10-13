package cache

import (
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Redis pool
var pool *redis.Pool

// buildDialer will build a redis connection from URL
func buildDialer(url string, options ...redis.DialOption) func() (redis.Conn, error) {
	return func() (redis.Conn, error) {
		return ConnectToURL(url, options...)
	}
}

// Connect creates a new connection pool connected to the specified url
func Connect(url string, maxActiveConnections, idleConnections, maxConnLifetime, idleTimeout int, dependencyMode bool, options ...redis.DialOption) (err error) {

	// Create a new pool
	pool = &redis.Pool{
		IdleTimeout:     time.Duration(idleTimeout) * time.Second,
		MaxActive:       maxActiveConnections,
		MaxConnLifetime: time.Duration(maxConnLifetime) * time.Second,
		MaxIdle:         idleConnections,
		Dial:            buildDialer(url, options...),
		TestOnBorrow: func(c redis.Conn, t time.Time) (err error) {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err = c.Do(pingCommand)
			return
		},
	}

	// Cleanup
	cleanUp()

	// Register scripts if enabled
	if dependencyMode {
		err = RegisterScripts()
	}

	return
}

// Disconnect closes the connection pool
func Disconnect() {
	if pool != nil {
		_ = pool.Close() //todo: handle this error?
	}
	pool = nil
}

// GetPool returns the underlying connection pool
func GetPool() *redis.Pool {
	return pool
}

// GetConnection will return a connection from the pool. The connection must be
// closed when done with use to return it to the pool
func GetConnection() redis.Conn {
	return pool.Get()
}

// ConnectToURL connects via REDIS_URL
// Source: github.com/soveran/redisurl
// URL Format: redis://localhost:6379
func ConnectToURL(urlString string, options ...redis.DialOption) (c redis.Conn, err error) {

	// Parse the URL
	var redisURL *url.URL
	if redisURL, err = url.Parse(urlString); err != nil {
		return
	}

	// Create the connection
	c, err = redis.Dial("tcp", redisURL.Host, options...)
	if err != nil {
		return
	}

	// Attempt authentication if needed
	if redisURL.User != nil {
		if password, ok := redisURL.User.Password(); ok {
			if _, err = c.Do("AUTH", password); err != nil {
				return
			}
		}
	}

	// Fire a select on DB
	if len(redisURL.Path) > 1 {
		db := strings.TrimPrefix(redisURL.Path, "/")
		_, err = c.Do("SELECT", db)
	}

	return
}

// cleanUp is fired after the pool is created
// Source: https://github.com/pete911/examples-redigo
func cleanUp() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		_ = pool.Close()
		os.Exit(0)
	}()
}
