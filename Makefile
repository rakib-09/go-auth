create_migration:
	@migrate create -ext sql -path database/migrations -seq $(NAME)
migrate_up:
	@migrate -path db/migrations -database "mysql://root:root@tcp(go-auth-mysql:3306)/go_db?multiStatements=true" -verbose up
migrate_down:
	@migrate -path db/migrations -database "mysql://root:root@tcp(go-auth-mysql:3306)/go_db?multiStatements=true" -verbose down


.PHONY: create_migration migrate_up migrate_down