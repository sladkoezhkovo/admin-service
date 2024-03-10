build:
	go build -v cmd/admin-service

debug:
	make build
	.\admin-service.exe -dotenv -config configs/local.yml

protoc:
	protoc proto/admin.proto --go_out=. --go-grpc_out=.

up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/psbd?sslmode=disable' up

down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/psbd?sslmode=disable' down