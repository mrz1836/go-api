#!/bin/bash

sqlboiler mysql --wipe && sed -i "" 's/fmt.Fprintf(tmp, "ssl-mode/\/\/fmt.Fprintf(tmp, "ssl-mode/' models/mysql_main_test.go