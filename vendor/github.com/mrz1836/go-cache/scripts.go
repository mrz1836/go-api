package cache

import "github.com/gomodule/redigo/redis"

// RegisterScripts register all scripts
func RegisterScripts() (err error) {

	// Register the kill dependency script
	killByDependencySha, err = RegisterScript(killByDependencyLua)

	// Return any error if found
	return
}

// RegisterScript register a script
func RegisterScript(script string) (sha string, err error) {

	// Get the connection
	c := GetConnection()

	// Defer close, skip error
	defer func() {
		_ = c.Close()
	}()

	// Set the script for killByDependency and return sha/error
	return redis.String(c.Do("SCRIPT", "LOAD", script))
}

// DidRegisterKillByDependencyScript returns true if the script has a sha from redis
func DidRegisterKillByDependencyScript() bool {
	return killByDependencySha != ""
}

// killByDependencySha is the SHA of the script
var killByDependencySha string

// killByDependencyLua is a script for kill related dependencies
var killByDependencyLua = `
--@begin=lua@
redis.replicate_commands()
local number_of_keys = table.getn(ARGV)
local all_keys = {}
for _, key in ipairs(ARGV) do
	table.insert(all_keys, key)
	local set = redis.call("SMEMBERS", key)
	for _, v in ipairs(set) do
	  table.insert(all_keys, v)
	end
end
return redis.call("DEL", unpack(all_keys))
--@end=lua@
`
