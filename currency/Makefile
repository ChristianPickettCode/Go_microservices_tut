.PHONY: protos

protos:
	export PATH="$PATH:$(go env GOPATH)/bin";
	protoc -I protos/ --go_out=protos/ --go-grpc_out=require_unimplemented_servers=false:protos/ protos/currency.proto