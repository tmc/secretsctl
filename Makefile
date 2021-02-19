API_COMMON_PROTOS?=$(shell go env GOPATH)/src/github.com/googleapis/api-common-protos
GOOGLEAPIS_PROTOS?=$(shell go env GOPATH)/src/github.com/googleapis/googleapis

.PHONY: protos
protos: deps
	buf protoc -I ${GOOGLEAPIS_PROTOS} -I . --go_gapic_out=./secrets --go_gapic_opt='go-gapic-package=github.com/tmc/secretsctl/secrets;secrets' google/cloud/secretmanager/v1/service.proto
	mv secrets/github.com/tmc/secretsctl/secrets/*.go ./secrets/
	buf protoc -I ${GOOGLEAPIS_PROTOS} -I . --go_cli_opt "root=secrets" --go_cli_opt "gapic=github.com/tmc/secretsctl/secrets" google/cloud/secretmanager/v1/service.proto --go_cli_out="cmd/secretsctl"

.PHONY: deps
deps:
	@command -v buf > /dev/null || go get github.com/bufbuild/buf/cmd/buf@latest
	@command -v protoc-gen-go_cli > /dev/null || go get github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_cli
	@command -v protoc-gen-go_gapic > /dev/null || go get github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic

.PHONY: clean
clean:
	rm -rf secrets/*
