/*
Package cache is a cache dependency management on-top of the famous redigo package
*/
package cache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Package constants (commands)
const (
	addToSetCommand      string = "SADD"
	deleteCommand        string = "DEL"
	dependencyPrefix     string = "depend:"
	evalCommand          string = "EVALSHA"
	executeCommand       string = "EXEC"
	existsCommand        string = "EXISTS"
	expireCommand        string = "EXPIRE"
	flushAllCommand      string = "FLUSHALL"
	getCommand           string = "GET"
	hashGetCommand       string = "HGET"
	hashMapGetCommand    string = "HMGET"
	hashKeySetCommand    string = "HSET"
	hashMapSetCommand    string = "HMSET"
	isMemberCommand      string = "SISMEMBER"
	keysCommand          string = "KEYS"
	listRangeCommand     string = "LRANGE"
	multiCommand         string = "MULTI"
	pingCommand          string = "PING"
	removeMemberCommand  string = "SREM"
	setCommand           string = "SET"
	setExpirationCommand string = "SETEX"
)

//GobCacheCreator can be used to generate the content to store. If the content
//is not found, the creator will be invoked and the result will be stored and
//returned
type GobCacheCreator func() ([]byte, error)

// Get gets a key from redis
func Get(key string) (string, error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the command
	return redis.String(conn.Do(getCommand, key))
}

// GetBytes gets a key from redis in bytes
func GetBytes(key string) ([]byte, error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the command
	return redis.Bytes(conn.Do(getCommand, key))
}

// GetStringSlice returns a []string stored in redis
func GetStringSlice(key string) (destination []string, err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// This command takes two parameters specifying the range: 0 start, -1 is the end of the list
	var values []interface{}
	values, err = redis.Values(conn.Do(listRangeCommand, key, 0, -1))
	if err != nil {
		return
	}

	// Scan slice by value, return with destination
	err = redis.ScanSlice(values, &destination)
	return
}

// GetAllKeys returns a []string of keys
func GetAllKeys() (keys []string, err error) {
	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Get all the keys
	return redis.Strings(conn.Do(keysCommand, "*"))
}

// Set will set the key in redis and keep a reference to each dependency
// value can be both a string or []byte
func Set(key string, value interface{}, dependencies ...string) (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the set command
	if _, err = conn.Do(setCommand, key, value); err != nil {
		return
	}

	// Link and return the error
	return linkDependencies(conn, key, dependencies...)
}

// SetExp will set the key in redis and keep a reference to each dependency
// value can be both a string or []byte
func SetExp(key string, value interface{}, ttl time.Duration, dependencies ...string) (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the set expiration
	_, err = conn.Do(setExpirationCommand, key, int64(ttl.Seconds()), value)
	if err != nil {
		return
	}

	// Link and return the error
	return linkDependencies(conn, key, dependencies...)
}

// Exists checks if a key is present or not
func Exists(key string) (bool, error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the command
	return redis.Bool(conn.Do(existsCommand, key))
}

// Expire sets the expiration for a given key
func Expire(key string, duration time.Duration) (err error) {
	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the expire command
	_, err = conn.Do(expireCommand, key, int64(duration.Seconds()))
	return
}

// Delete is an alias for KillByDependency()
func Delete(keys ...string) (total int, err error) {
	return KillByDependency(keys...)
}

// HashSet will set the hashKey to the value in the specified hashName and link a
// reference to each dependency for the entire hash
func HashSet(hashName, hashKey string, value interface{}, dependencies ...string) (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Set the hash key
	if _, err = conn.Do(hashKeySetCommand, hashName, hashKey, value); err != nil {
		return
	}

	// Link and return the error
	return linkDependencies(conn, hashName, dependencies...)
}

// HashGet gets a key from redis via hash
func HashGet(hash, key string) (string, error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the command
	return redis.String(conn.Do(hashGetCommand, hash, key))
}

// HashMapGet gets values from a hash map for corresponding keys
func HashMapGet(hashName string, keys ...interface{}) ([]string, error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Build up the arguments
	keys = append([]interface{}{hashName}, keys...)

	// Fire the command with all keys
	return redis.Strings(conn.Do(hashMapGetCommand, keys...))
}

// HashMapSet will set the hashKey to the value in the specified hashName and link a
// reference to each dependency for the entire hash
func HashMapSet(hashName string, pairs [][2]interface{}, dependencies ...string) (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Set the arguments
	args := make([]interface{}, 0, 2*len(pairs)+1)
	args = append(args, hashName)
	for _, pair := range pairs {
		args = append(args, pair[0])
		args = append(args, pair[1])
	}

	// Set the hash map
	if _, err = conn.Do(hashMapSetCommand, args...); err != nil {
		return
	}

	// Link and return the error
	return linkDependencies(conn, hashName, dependencies...)
}

// HashMapSetExp will set the hashKey to the value in the specified hashName and link a
// reference to each dependency for the entire hash
func HashMapSetExp(hashName string, pairs [][2]interface{}, ttl time.Duration, dependencies ...string) (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Set the arguments
	args := make([]interface{}, 0, 2*len(pairs)+1)
	args = append(args, hashName)
	for _, pair := range pairs {
		args = append(args, pair[0])
		args = append(args, pair[1])
	}

	// Set the hash map
	if _, err = conn.Do(hashMapSetCommand, args...); err != nil {
		return
	}

	// Fire the expire command
	if _, err = conn.Do(expireCommand, hashName, ttl.Seconds()); err != nil {
		return
	}

	// Link and return the error
	return linkDependencies(conn, hashName, dependencies...)
}

// SetAdd will add the member to the Set and link a reference to each dependency
// for the entire Set
func SetAdd(setName, member interface{}, dependencies ...string) (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Add member to set
	if _, err = conn.Do(addToSetCommand, setName, member); err != nil {
		return
	}

	// Link and return the error
	return linkDependencies(conn, setName, dependencies...)
}

// SetIsMember returns if the member is part of the set
func SetIsMember(set, member interface{}) (bool, error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Check if is member
	return redis.Bool(conn.Do(isMemberCommand, set, member))
}

// SetRemoveMember removes the member from the set
func SetRemoveMember(set, member interface{}) (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Remove and return
	_, err = conn.Do(removeMemberCommand, set, member)
	return
}

// GetOrSetWithExpirationGob will return the cached value for the key or use the
// GobCacheCreator to create and insert the value into the cache. If the expiration
// time is set to a value greater than zero, the key will be set to expire in
// the provided duration
func GetOrSetWithExpirationGob(key string, fn GobCacheCreator, duration time.Duration, dependencies ...string) (data []byte, err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	//Get from redis
	data, err = redis.Bytes(conn.Do(getCommand, key))

	//Set the string in redis
	if err != nil {

		//Set the value
		data, err = fn()
		if err != nil {
			return
		}

		//No data?!
		if len(data) == 0 {
			err = fmt.Errorf("value is empty for key: %s", key)
			return
		}

		//Go routine to set the key and expiration
		go func(key string, data []byte, duration time.Duration, dependencies []string) {

			// Create a new connection and defer closing
			conn := GetConnection()
			defer func() {
				_ = conn.Close()
			}()

			//Set an expiration time if found
			if duration > 0 {
				_ = SetExp(key, data, duration, dependencies...) //todo: handle the error?
			} else {
				_ = Set(key, data, dependencies...) //todo: handle the error?
			}
		}(key, data, duration, dependencies)
	}

	//Return the value
	return
}

// DestroyCache will flush the entire redis server. It only removes keys, not
// scripts
func DestroyCache() (err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Fire the command
	_, err = conn.Do(flushAllCommand)
	return
}

// KillByDependency removes all keys which are listed as depending on the key(s)
// Also: Delete()
func KillByDependency(keys ...string) (total int, err error) {

	// Create a new connection and defer closing
	conn := GetConnection()
	defer func() {
		_ = conn.Close()
	}()

	// Do we have keys to kill?
	if len(keys) == 0 {
		return
	}

	// Create the arguments
	args := make([]interface{}, len(keys)+2)
	deleteArgs := make([]interface{}, len(keys))

	args[0] = killByDependencySha
	args[1] = 0

	// Loop keys
	for i, key := range keys {
		args[i+2] = dependencyPrefix + key
		deleteArgs[i] = key
	}

	// Create the script
	total, err = redis.Int(conn.Do(evalCommand, args...))
	if err != nil {
		return
	}

	// Fire the delete
	_, err = conn.Do(deleteCommand, deleteArgs...)
	return
}

// linkDependencies links any dependencies
func linkDependencies(conn redis.Conn, key interface{}, dependencies ...string) (err error) {

	// No dependencies given
	if len(dependencies) == 0 {
		return
	}

	// Send the multi command
	if err = conn.Send(multiCommand); err != nil {
		return
	}

	// Add all to the set
	for _, dependency := range dependencies {
		if err = conn.Send(addToSetCommand, dependencyPrefix+dependency, key); err != nil {
			return
		}
	}

	// Fire the exec command
	_, err = redis.Values(conn.Do(executeCommand))
	return
}
