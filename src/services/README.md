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
