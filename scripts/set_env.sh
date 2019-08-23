#!/bin/bash

# Port for the API http requests
export API_SERVER_PORT=3000

# Environment to run
export API_ENVIRONMENT=development

#
# Read Database
#
export API_DATABASE_READ__HOST=localhost
export API_DATABASE_READ__USER=apiDbTestUser
export API_DATABASE_READ__PASSWORD=ThisIsSecureEnough123

#
# Write Database
#
export API_DATABASE_WRITE__HOST=localhost
export API_DATABASE_WRITE__USER=apiDbTestUser
export API_DATABASE_WRITE__PASSWORD=ThisIsSecureEnough123

# If you want to use cache, set the redis url
# Example: redis://localhost:6379
export API_CACHE__URL='redis://localhost:6379'