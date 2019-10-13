# Drop the current database
DROP DATABASE IF EXISTS api_example;

# Build a new database
CREATE DATABASE api_example;

# Drop any existing user
DROP USER IF EXISTS 'apiDbTestUser'@'%';

# Set the SQL mode
SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

# Create the user and assign grants
CREATE USER 'apiDbTestUser'@'%' IDENTIFIED BY 'ThisIsSecureEnough123';
GRANT USAGE ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT DELETE ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT INSERT ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT SELECT ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT UPDATE ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT CREATE ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT INDEX ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT ALTER ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT REFERENCES ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT DROP ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT LOCK TABLES ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT CREATE TEMPORARY TABLES ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT TRIGGER ON `api_example`.* TO 'apiDbTestUser'@'%';
GRANT SUPER ON *.* TO 'apiDbTestUser'@'%';
FLUSH PRIVILEGES;

# Database was generated from SQLBoiler
GRANT ALL ON `fqgmsvujyikggnnkpkfbldubvmtcirjhohkivzex`.* to 'apiDbTestUser'@'%';
GRANT ALL ON `evczgdvdwlgpajxgqflurknmyvlxzpksrermjoix`.* to 'apiDbTestUser'@'%';