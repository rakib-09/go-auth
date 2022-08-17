create_migration:
	migrate create -ext sql -path database/migrations -seq $(NAME)
migrate_up:
	migrate -path database/migrations -database "mysql://root:root@tcp(localhost:3306)/go_db?multiStatements=true" -verbose up
migrate_down:
	migrate -path database/migrations -database "mysql://root:root@tcp(localhost:3306)/go_db?multiStatements=true" -verbose down


.PHONY: create_migration migrate_up migrate_down