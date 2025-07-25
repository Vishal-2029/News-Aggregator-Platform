build:
	go build -o News Aggregator Platform main.go

run: 
	go run main.go

restapi:
	go run cmd/REST-Api/rest.go

tidy:
	go mod tidy