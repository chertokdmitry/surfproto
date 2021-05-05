gen:
	protoc --proto_path=src/proto src/proto/*.proto --go_out=plugins=grpc:src/pb

clean:
	rm src/pb/*.go

server:
	go run main.go