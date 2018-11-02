regenerate:
	protoc --proto_path=$(go env GOPATH)/src --proto_path . --go_out=plugins=grpc:. types.proto
