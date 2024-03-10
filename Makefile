build:
	go build -v cmd/app

protoc:
	protoc proto/admin.proto --go_out=. --go-grpc_out=.