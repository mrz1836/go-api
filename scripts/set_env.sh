#!/bin/bash

# If you want to use cache, set the redis url
# Example: redis://localhost:6379
export CACHE_URL=your-cache-url

#
# Read Database
#
export DATABASE_DRIVER=mysql
export DATABASE_MAX_CONN_LIFETIME=60
export DATABASE_READ_HOST=locahost
export DATABASE_READ_MAX_IDLE_CONNECTIONS=5
export DATABASE_READ_MAX_OPEN_CONNECTIONS=5
export DATABASE_READ_NAME=api_example
export DATABASE_READ_PORT=3306
export DATABASE_READ_USER=apiDbTestUser
export DATABASE_READ_PASSWORD=ThisIsSecureEnough123

#
# Write Database
#
export DATABASE_DRIVER=mysql
export DATABASE_MAX_CONN_LIFETIME=60
export DATABASE_WRITE_HOST=locahost
export DATABASE_WRITE_MAX_IDLE_CONNECTIONS=5
export DATABASE_WRITE_MAX_OPEN_CONNECTIONS=5
export DATABASE_WRITE_NAME=api_example
export DATABASE_WRITE_PORT=3306
export DATABASE_WRITE_USER=apiDbTestUser
export DATABASE_WRITE_PASSWORD=ThisIsSecureEnough123