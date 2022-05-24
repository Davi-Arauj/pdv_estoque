postgres:
	docker run --name gomaster -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
createdb:
	docker exec -it gomaster createdb --username=root --owner=root pdv_estoque
dropdb:
	docker exec -it gomaster dropdb pdv_estoque
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/pdv_estoque?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/pdv_estoque?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY:postgres createdb dropdb migrateup migratedown sqlc