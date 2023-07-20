build:
	go build -o todolist cmd/main.go

run:
	go run cmd/main.go

docker:
	docker build --tag todolist .
	docker run -e TODOLIST_CONFIG_PATH=./config/config.yaml --network=host todolist