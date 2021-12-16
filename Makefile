gogen:
	go generate ./...

gotest: gogen
	go test -v ./...

build:
	docker build --tag todo-app:latest .

run: build
	docker-compose -f docker-compose.yaml -p todo-app --env-file .env up
