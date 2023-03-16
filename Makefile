LOCAL_BIN:=$(CURDIR)/bin

run:
	go run cmd/events-api/main.go

lint:
	golangci-lint run

test:
	go test -v ./...

# Build mod
deps:
		ls go.mod || go mod init github.com/Vovchikus/events-api
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u google.golang.org/protobuf
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get google.golang.org/grpc
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/envoyproxy/protoc-gen-validate
		go get -u gopkg.in/yaml.v3
		go get -u go.uber.org/zap
		go get -u github.com/Shopify/sarama
		go get github.com/rs/zerolog/log
		go get golang.org/x/sync
		go get github.com/sarulabs/di
		go get github.com/allegro/bigcache
		go get github.com/olivere/elastic/v7
		go get github.com/tamerh/xml-stream-parser
		go get github.com/schollz/progressbar/v3
		go get github.com/sirupsen/logrus
		go get gopkg.in/natefinch/lumberjack.v2
		go get gopkg.in/jeevatkm/go-model.v1
		go get github.com/paulmach/osm
		go get github.com/dustin/go-humanize
		go get github.com/spf13/viper
		go get github.com/ProsviryakovVadim/bigcache-lib
		go get github.com/gojuno/minimock/v3
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate
		go mod tidy -go=1.20
		#go mod tidy -go=1.16 && go mod tidy -go=1.17
		#go mod tidy -compat=1.17

# Build project
.PHONY: build
build: .vendor-proto .generate .bin

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/events-api
		yes | cp -rf api/events-api/* vendor.protogen/api/events-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
  			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/grpc-ecosystem &&\
			mkdir -p vendor.protogen/protoc-gen-openapiv2 &&\
			mv vendor.protogen/grpc-ecosystem/protoc-gen-openapiv2 vendor.protogen &&\
			rm -rf vendor.protogen/grpc-ecosystem ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/events-api
		protoc -I vendor.protogen \
				--go_out=pkg/events-api --go_opt=paths=import \
				--go-grpc_out=pkg/events-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/events-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/events-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/events-api/v1/events-api.proto
		mv pkg/events-api/github.com/Vovchikus/events-api/pkg/events-api/* pkg/events-api/
		rm -rf pkg/events-api/github.com
		mkdir -p cmd/events-api
		cd pkg/events-api && ls go.mod || go mod init github.com/Vovchikus/events-api/pkg/events-api && go mod tidy

.PHONY: .bin
.bin:
		CGO_ENABLED=0 GOOS=linux go build -o bin/events-api cmd/events-api/*.go

# Start docker containers
.PHONY: start
start: .compose-build .compose-up

.PHONY: .compose-build
.compose-build:
	docker-compose build

.PHONY: .compose-up
.compose-up:
	docker-compose up -d

# Stop docker containers
.PHONY: stop
stop:
	docker-compose stop

# MiniMock
MODULE_NAME=github.com/Vovchikus/events-api
MINI_MOCK_TAG:=3.0.8
LOCAL_BI:=$(CURDIR)/bin
MINI_MOCK_BIN=$(LOCAL_BIN)/minimock

.PHONY: .install-third-party-dependencies
.install-third-party-dependencies:
	$(info Downloading third party dependencies)
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && \
		GOBIN=$(LOCAL_BIN) go get github.com/gojuno/minimock/v3/cmd/minimock@v$(MINI_MOCK_TAG) && \
	rm -rf $$tmp

.PHONY: bin-deps
bin-deps: .install-third-party-dependencies
