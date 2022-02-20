.PHONY: rundb createdb dropdb migrateup migratedown sqlcgen 

runDBContainer:
	docker run  --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

removeDBContainer:
	docker rm postgres12

createDB:
	docker exec -it postgres12 createdb --username=root --owner=root simple-bank

dropDB:
	docker exec -it postgres12 dropdb simple-bank

migrateUp:
	migrate -path store/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose up

migrateDown:
	migrate -path store/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose down

sqlcGen:
	sqlc generate

connectPSQL:
	psql -Atx "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable"

runKube:
		kubectl ns duckhue01 
		kubectl apply -f k8s/deployment.yml 
		kubectl apply -f k8s/service.yml 
		kubectl apply -f k8s/ingress.yml 
		kubectl apply -f k8s/issuer.yml