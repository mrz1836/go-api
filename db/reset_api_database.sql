# Drop the current database
DROP DATABASE IF EXISTS api_example;

# Build a new database
CREATE DATABASE api_example;

# Drop any existing user
DROP USER IF EXISTS 'apiDbTestUser'@'%';

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