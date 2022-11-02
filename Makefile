start:
	go run main.go

test:
	go test -v ./tests

mocks :
	mockgen -source=services/repository.go -destination=services/repository/postgres/mocks/repository.go -package=mocks

serve:
	nodemon --exec go run main.go --signal SIGTERM