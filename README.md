# training-golang-project-
This project use to training golang 

-setup go path: export GOPATH=$HOME/go
                export PATH=$PATH:$(go env GOPATH)/bin

go mod init 
go mod tidy
go mod vendor



Object: 
  - Product(id, name, price, description)
  - Order(id, products, orderBy)
  - User(id, name, email)

#### download and install go migrate:
  - link: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
  - go get github.com/golang-migrate/migrate/v4  

##### generate db/migrations file
migrate create -ext sql -dir [path_to/migrations] [name_file]  


# database local
AWSTIO_DB_URL=postgresql://vietle:123456@localhost:5432/awstdev?sslmode=disable

migrate -path ./migrations -database "postgresql://vietle:123456@localhost:5432/awstdev?sslmode=disable" -verbose up

# run test
go test -cover -v  ./infrastructure/dao/datastoredao -run UpdateCurrentCreatorProfile
go test -v -cover -run TestCreateUser