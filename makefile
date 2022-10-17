dev:
	make build
	make run

build:
	go build -o ./bin/main ./cmd/main.go

run:
	./bin/main run --config=./config.yml
	