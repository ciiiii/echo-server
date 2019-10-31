version := 0.1.0

proto:
	@protoc --proto_path=. --go_out=plugins=grpc:grpcServer echo.proto
run:
	@go run main.go
ping:
	@grpcurl --plaintext localhost:8000 grpcServer.Echo.Ping
build:
	@docker build -t killercai/grpc-echo:$(version) .
push:
	@dockr push killercai/grpc-echo:$(version)