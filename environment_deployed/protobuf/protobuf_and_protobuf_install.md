# Istall protobuf on macOS
[the response of protobuf on the github](https://github.com/protocolbuffers/protobuf/releases)

## install protobuf by homebrew
> brew install protobuf

## checkout protobuf is installed 
> protoc --version

## export `GOBIN` && `environment path `on the `~/.zshrc`
```bash
vim ~/.zshrc

export GOBIN=$GOPATH/bin
export PATH="$GOBIN:$PATH"
```
ps: ensure you are config `gopath` first.
## install `protoc-gen-go`
> go get -u -v google.golang.org/grpc  
> go get -v -u github.com/golang/protobuf/protoc-gen-go

## test
you can checkout the grpc and protobuf is installed by flowing the test example:

### general pb.go 
> protoc --go_out=plugins=grpc:. helloworld.proto

### run the test program
> go run main.go
```bash
2022/07/27 17:32:42 server listening at [::]:50051
2022/07/27 17:32:43 Received: world
2022/07/27 17:32:43 Greeting: Hello world
```
