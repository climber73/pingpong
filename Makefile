regenerate:
	protoc --proto_path=$(GOPATH)/src --proto_path . --go_out=plugins=grpc:. types/types.proto
