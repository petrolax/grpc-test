all: hello bye
	
	
hello:
	protoc --go_out=. --go-grpc_out=. \
	config/hello.proto

bye:
	protoc --go_out=. --go-grpc_out=. \
	config/bye.proto