postgres:
	docker run --name simplebank-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
createdb:
	docker exec -it simplebank-db createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it simplebank-db dropdb -U postgres simple_bank
migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./db/sqlc
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/althof3/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock
