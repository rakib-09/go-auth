#! /usr/bin/env bash

# Let the DB start
sleep 5;
# Run migration
make migrate_up
#migrate -path db/migrations -database "mysql://root:root@tcp(go-auth-mysql:3306)/go_db?multiStatements=true" -verbose up
# Run application using air
air