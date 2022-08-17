migrateup:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/dev?sslmode=disable" -verbose up

migratedown:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/dev?sslmode=disable" -verbose down 1

migrateforce:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/dev?sslmode=disable" force 2


# orm volatiletech/sqlboiler mapping at local
generate-models:
	sqlboiler psql --config sqlboiler.toml


# mockery
generate-mock:
	GO111MODULE=on go get github.com/vektra/mockery/v2/.../
	mockery --dir ./internal/databases/product/ --all --keeptree