protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. proto/helloworld.proto

go build -buildmode=plugin -o plugin/plugin.so ./plugin

#PATH=$PATH:$GOPATH:$GOBIN
#export PATH