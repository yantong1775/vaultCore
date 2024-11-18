postgres:
	docker run --name mypostgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d -p 5432:5432 postgres

createdb:
	docker exec -it mypostgres createdb --username=root --owner=root vaultCore

dropdb:
	docker exec -it mypostgres dropdb vaultCore

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/vaultCore?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/vaultCore?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
.PHONY: postgres createdb dropdb sqlc migrateup migratedown test server 