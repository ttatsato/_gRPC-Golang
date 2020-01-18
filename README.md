# macでprotocコマンドを使う
```shell script
brew instal protobuf
```

## .proto ファイルからserver、client,interface等のコードを生成する
```shell script
# .proto ファイルからserver、client,interface等のコードを生成する
protoc --go_out=plugins=grpc:../pb sample.proto
```

## protoc docの生成
### プラグインのインストール
```shell script
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
```

### doc生成
```shell script
protoc --doc_out=html,index.html:./ proto/*.proto
```