#!/bin/bash

# Runs locally to reset your database and seed the new database
mysql -u root < ./db/reset_api_database.sql && goose -dir "db/migrations" mysql "apiDbTestUser:ThisIsSecureEnough123@/api_example?parseTime=true" up