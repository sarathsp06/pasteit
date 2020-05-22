GO=go 
proto:
	~/bin/protoc \
		--proto_path=${GOPATH}/src:.\
		--twirp_out=. \
	 	--go_out=. \
		./api/protobuf/pasteit.proto 

lint:
	golangci-lint run ./...

bin:
	mkdir -p bin

build: bin
	${GO} build -o ./bin/server ./cmd/server