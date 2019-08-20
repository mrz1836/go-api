#!/bin/bash

# Port for the API http requests
export API_SERVER_PORT=3000

# Environment to run
export API_ENVIRONMENT=development

#
# Read Database
#
export API_DATABASE_DRIVER=mysql
export API_DATABASE_MAX_CONN_LIFETIME=60
export API_DATABASE_READ_HOST=locahost
export API_DATABASE_READ_MAX_IDLE_CONNECTIONS=5
export API_DATABASE_READ_MAX_OPEN_CONNECTIONS=5
export API_DATABASE_READ_NAME=api_example
export API_DATABASE_READ_PORT=3306
export API_DATABASE_READ_USER=apiDbTestUser
export API_DATABASE_READ_PASSWORD=ThisIsSecureEnough123

#
# Write Database
#
export API_DATABASE_DRIVER=mysql
export API_DATABASE_MAX_CONN_LIFETIME=60
export API_DATABASE_WRITE_HOST=locahost
export API_DATABASE_WRITE_MAX_IDLE_CONNECTIONS=5
export API_DATABASE_WRITE_MAX_OPEN_CONNECTIONS=5
export API_DATABASE_WRITE_NAME=api_example
export API_DATABASE_WRITE_PORT=3306
export API_DATABASE_WRITE_USER=apiDbTestUser
export API_DATABASE_WRITE_PASSWORD=ThisIsSecureEnough123

# If you want to use cache, set the redis url
# Example: redis://localhost:6379
export API_CACHE_URL=your-cache-url