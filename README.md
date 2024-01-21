# gogrpc


## ProtoBuf

* 1. proto 파일 작성
```protobuf
syntax = "proto3";

package greet;

option go_package = "github.com/jc3wrld999/gogrpc/greet/proto";

message Dummy {
  uint32 id = 1;
}

service DummyService {

}
```
* 2. 컴파일
```bash
protoc greet/proto/greet.proto --go_out=. --go_opt=module=github.com/jc3wrld999/gogrpc --go-grpc_out=. --go-grpc_opt=module=github.com/jc3wrld999/gogrpc/greet/proto
```

- unary
