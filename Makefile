ip ?= 10000
client:
	go run main.go client
normal-node:
	go run main.go node --ip $(ip)
byzantine-node:
	go run main.go node --ip $(ip) --byzantine
proto:
	protoc --proto_path=message --go_out=paths=source_relative:message --go-grpc_out=paths=source_relative:message message/*.proto
