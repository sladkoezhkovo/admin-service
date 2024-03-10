build:
	go build -v cmd/admin-service

debug:
	make build
	.\admin-service.exe -dotenv -config configs/local.yml

protoc:
	protoc proto/admin.proto --go_out=. --go-grpc_out=.