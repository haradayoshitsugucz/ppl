# Go parameters
GO_ENV=GO111MODULE=on GOPRIVATE="github.com/haradayoshitsugucz"
GOBUILD=CGO_ENABLED=0 $(GO_ENV) go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"'
BINARY_NAME=purple-api

# ビルド
build: wire_gen tidy
	$(GOBUILD) -o $(BINARY_NAME) -v

# ローカル環境起動
run_local: build
	./$(BINARY_NAME) -env=local

update:
	go get -u -v ./...

tidy:
	go mod tidy

wire_gen:
	cd di/ && wire

test_e2e:
	docker-compose -f docker/test/docker-compose.yml up -d
	go clean -testcache
	go test -race -v ./presentation/router
	docker-compose -f docker/test/docker-compose.yml down

test_model:
	go clean -testcache
	go test -race -v ./presentation/parameter
	go test -race -v ../domain/model
