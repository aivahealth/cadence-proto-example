all: protos
	go build -o example *.go

protos:
	protoc --go_out=. example.proto

