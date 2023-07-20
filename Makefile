protoAuth:
	protoc --go_out=. --go-grpc_out=. pkg/auth/pb/auth.proto
protoUser:
	protoc --go_out=. --go-grpc_out=. pkg/user/pb/user.proto

run:
	go run cmd/api/main.go