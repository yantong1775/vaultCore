postgres:
	docker run --name mypostgres --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d -p 5432:5432 postgres

createdb:
	docker exec -it mypostgres createdb --username=root --owner=root vaultCore

dropdb:
	docker exec -it mypostgres dropdb vaultCore

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/vaultCore?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/vaultCore?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/vaultCore?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/vaultCore?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/yantong1775/vaultCore/db/sqlc Store

.PHONY: postgres createdb dropdb sqlc migrateup migrateup1 migratedown migratedown1 test server mock mock