GO=go 
GCLOUD=gcloud
GCP_PROJECT=leadmrktr
APP=pasteit

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

build_cloud: 
	${GCLOUD} builds submit --tag gcr.io/${GCP_PROJECT}/${APP}
	
release_cloudrun:
	${GCLOUD} run deploy --image gcr.io/${GCP_PROJECT}/${APP} --platform managed