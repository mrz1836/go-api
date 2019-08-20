#!/bin/bash

# Wipes out the folder, replace the ssl-mode (fix for MariaDB)
sqlboiler mysql --wipe && sed -i "" 's/fmt.Fprintf(tmp, "ssl-mode/\/\/fmt.Fprintf(tmp, "ssl-mode/' ./models/generated/mysql_main_test.go