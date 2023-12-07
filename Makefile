postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root go_finance

dropdb:
	docker exec -it postgres14 dropdb go_finance	

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_finance?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb dropdb migrateup migrationdrop sqlc test server