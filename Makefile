API_COMMON_PROTOS?=$(shell go env GOPATH)/src/github.com/googleapis/api-common-protos
GOOGLEAPIS_PROTOS?=$(shell go env GOPATH)/src/github.com/googleapis/googleapis

protos:
	rm -rf secrets/*
	protoc -I ${API_COMMON_PROTOS} -I ${GOOGLEAPIS_PROTOS} --go_gapic_out ./secrets --go_gapic_opt 'go-gapic-package=github.com/tmc/secretsctl/secrets;secrets' google/cloud/secretmanager/v1/service.proto
	mv secrets/github.com/tmc/secretsctl/secrets/*.go ./secrets/
	protoc -I ${GOOGLEAPIS_PROTOS} --go_cli_out=cmd/secretsctl --go_cli_opt "root=secrets" --go_cli_opt "gapic=github.com/tmc/secretsctl/secrets" google/cloud/secretmanager/v1/service.proto
