# このディレクトリの役割

`routes`から要求された DB 操作などを行い,結果を返す

行う操作の内容としては以下のようなものを想定している

- Servant への gRPC を用いた制御命令の送信
- DB 操作

## gRPC について

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

#### protoc

プロトコル定義ファイルから各言語のクラスを生成する protoc というツールが用意されているのでそれを使う.

インストール方法は [How can I install protoc on Ubuntu 16.04?](https://askubuntu.com/questions/1072683/how-can-i-install-protoc-on-ubuntu-16-04) を参照.
