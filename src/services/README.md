# このディレクトリの役割

`routes`から要求された DB 操作などを行い,結果を返す

行う操作の内容としては以下のようなものを想定している

- Servant への gRPC を用いた制御命令の送信
- DB 操作

## gRPC

### gRPC とはなんぞや

gRPC の説明は[さくらインターネットの記事](https://knowledge.sakura.ad.jp/24059/)がわかりやすい.

protocol buffer language については [公式ドキュメント](https://developers.google.com/protocol-buffers/docs/proto3) を参照.

- [サービス間通信のための新技術「gRPC」入門](https://knowledge.sakura.ad.jp/24059/)
- [Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3#adding_comments)

### gRPC や関連ツールのインストール

#### gRPC ライブラリ

各言語 gRPC ライブラリが用意されている.

今回は Python と Go を使うのでこの 2 つの言語の gRPC ライブラリのインストールについてのみ書く

##### Python

```bash
pip3 install grpcio
```

[grpcio](https://github.com/grpc/grpc)

##### Go

```bash
go get -u google.golang.org/grpc
```

[gRPC-go](https://github.com/grpc/grpc-go)

### .proto からコードを生成する

プロトコル定義ファイルから各言語のクラスを生成する protoc というツールが用意されているのでそれを使う.

インストール方法は [How can I install protoc on Ubuntu 16.04?](https://askubuntu.com/questions/1072683/how-can-i-install-protoc-on-ubuntu-16-04) を参照.

#### GO

[Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)

## 参考にしたサイト

- [Can I define a grpc call with a null request or response?](https://stackoverflow.com/questions/31768665/can-i-define-a-grpc-call-with-a-null-request-or-response): gRPC の rpc で Null のような値を扱う方法. rpc では必ずメッセージを指定する必要があるので, Empty というメッセージを定義して, それを引数と返り値にする.
- [How to return an array in Protobuf service rpc](https://stackoverflow.com/questions/43167762/how-to-return-an-array-in-protobuf-service-rpc): rpc の returns の値に配列を入れる方法. 配列は使えず, stream 型を使うか, 新たなメッセージを定義する必要がある.
- [gRPC(Go) で API を実装する](https://blog.fenrir-inc.com/jp/2016/10/grpc-go.html)
- [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
