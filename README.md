# training-golang-project
  This project use to training golang. Project describes the process of customers order products.


# How to setup gopath:
  - export GOPATH=$HOME/go
  - export PATH=$PATH:$(go env GOPATH)/bin


# download and install go migrate:
  - link: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
  - go get github.com/golang-migrate/migrate/v4


# how to setup sqlboilder:
  - please read document: https://github.com/volatiletech/sqlboiler    


##### generate db/migrations file
  - migrate create -ext sql -dir [path_to/migrations] [name_file]  


##### How to run project at local
  - export variables in file env.dev
  - in terminal: go run cmd/handlers/main.go

##### How to run project with Docker
  - in terminal: docker compose up 
  - open other terminal(to check containers running): docker ps

##### run test
  - Just run a func test: go test -cover -v [path/to/folder] -run [func_test_name]
  - Run all tests in project: go test ./...
