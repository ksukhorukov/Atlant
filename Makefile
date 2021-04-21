compile:
	#go get -v
	protoc api/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

	go build -o ./server/server ./server/server.go
	go build -o ./client/client ./client/client.go