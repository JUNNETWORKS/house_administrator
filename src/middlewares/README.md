# このディレクトリの役割

gin のリクエストに対する処理の流れは以下のようなものである.

```plain
Request -> Route Parser -> Middleware -> Route Handler -> Middleware -> Response
```

実際処理する関数の前に必ず Middleware は通る形になる.

ミドルウェアの使用例

- アクセスログの記録
- ユーザー認証
- セッションの管理など

TODO:

フロントからの API のリクエスト URL は`api/<room_name>/hogehoge`のようになる. そこで, ミドルウェアで `<room_name>` から 接続先のローカル IP に変換する処理を実装することで, controller の中で部屋と IP アドレスの変換処理を行わなくても良いようにする. [Pass var to route handler #420](https://github.com/gin-gonic/gin/issues/420)
