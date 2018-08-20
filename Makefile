GOCMD=go
GOMOD=$(GOCMD) mod
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

API_BINARY_NAME=bin/api
API_BINARY_UNIX=$(CTL_BINARY_NAME)_unix
CTL_BINARY_NAME=bin/ctl
CTL_BINARY_UNIX=$(CTL_BINARY_NAME)_unix

all: test build
build: build-api build-ctl
build-api:
	$(GOBUILD) -o $(API_BINARY_NAME) -v ./cmd/api
build-ctl:
	$(GOBUILD) -o $(CTL_BINARY_NAME) -v ./cmd/ctl
test: 
	$(GOTEST) -v ./...
clean: clean-api clean-ctl
clean-api: 
	$(GOCLEAN) ./cmd/api
	rm -f $(API_BINARY_NAME)
	rm -f $(API_BINARY_UNIX)
clean-ctl: 
	$(GOCLEAN) ./cmd/ctl
	rm -f $(CTL_BINARY_NAME)
	rm -f $(CTL_BINARY_UNIX)
run-api:
	$(GOBUILD) -o $(API_BINARY_NAME) -v ./cmd/api/...
	./$(API_BINARY_NAME)
run-ctl:
	$(GOBUILD) -o $(CTL_BINARY_NAME) -v ./cmd/ctl/...
	./$(CTL_BINARY_NAME)

# Cross compilation
build-api-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(API_BINARY_UNIX) -v ./cmd/api
docker-api-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/powerssl.io golang:1.11rc1 go build -o "$(API_BINARY_UNIX)" -v ./cmd/api
build-ctl-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(CTL_BINARY_UNIX) -v ./cmd/ctl
docker-ctl-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/powerssl.io golang:1.11rc1 go build -o "$(CTL_BINARY_UNIX)" -v ./cmd/ctl

generate:
	protoc \
		--proto_path=. \
		--proto_path=$$($(GOMOD) download -json github.com/gogo/protobuf | jq -r '.Dir') \
		--proto_path=$$($(GOMOD) download -json github.com/gogo/googleapis | jq -r '.Dir') \
		--gogo_out=paths=source_relative,plugins=grpc,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:.\
		api/v1/*.proto
