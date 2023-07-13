proto:
	protoc --go_out=. --go-grpc_out=. pkg/auth/pb/auth.proto