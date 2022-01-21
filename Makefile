.PHONY: rundb createdb dropdb migrateup migratedown sqlcgen test

runDBContainer:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

removeDBContainer:
	docker rm postgres12

createDB:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropDB:
	docker exec -it postgres12 dropdb simple_bank

migrateUp:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlcGen:
	sqlc generate

start:
	make runDBContainer
	make createDB
	make migrateup