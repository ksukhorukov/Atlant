compile:
# 	export GO111MODULE=on
# 	go get google.golang.org/grpc
#		go get go.mongodb.org/mongo-driver
#		go get github.com/gabriel-vasile/mimetype
# 	go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
# 	export PATH="$PATH:$(go env GOPATH)/bin"

	#go get -v
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./proto/ --go-grpc_opt=paths=source_relative proto/*.proto
	go build