version := 0.1.0

run:
	@go run main.go
build:
	@docker build -t killercai/tcp-echo:$(version) .
push:
	@docker push killercai/tcp-echo:$(version)