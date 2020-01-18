# macでprotocコマンドを使う
```shell script
brew instal protobuf
```

## protocコマンド実行場所
```shell script
mkdir proto
mkdir proto/doc
mkdir proto/pb
cd proto
```

## .protoファイルにインターフェースを定義する
```proto
// sample.proto
syntax = "proto3";
service Cat {
  rpc GetMyCat (GetMyCatMessage) returns (MyCatResponse) {}
}
message GetMyCatMessage {
  string target_cat = 1;
}
message MyCatResponse {
  string name = 1;
  string kind = 2;
}
```

## .proto ファイルからserver、client,interface等のコードを生成する
```shell script
# .proto ファイルからserver、client,interface等のコードを生成する
protoc --go_out=plugins=grpc:./pb sample.proto
```

## protoc docの生成
### プラグインのインストール
```shell script
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
```

### doc生成
```shell script
protoc --doc_out=html,index.html:./doc *.proto
```