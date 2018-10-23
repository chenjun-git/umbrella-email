all: pb

pb: *.proto
	protoc -I/Users/chenjun/umbrella/src/umbrella-common/proto --proto_path=$(GOPATH)/src:. --go_out=plugins=grpc:. *.proto

.PHONY: pb
