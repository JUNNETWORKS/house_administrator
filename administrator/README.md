# 各ディレクトリの役割

- `backend/`: servant やフロントエンドアプリケーションからくるリクエストに応じて DB との通信や servant への命令などを行う Administrator のコア部分
  - `main.go`: ルーティングやミドルウェアの設定などを行っている
  - `data/`: 各種構造体や DB アクセスなどの関数が入ってる
  - `handlers/`: 各 URL へのアクセスに対する処理を行う
  - `utils/`: 便利な汎用的に使える関数とかが入ってる
- `db/`: バックエンド DB の初期化用の SQL ファイルが入ってる
- `front/`: フロントエンドプロジェクト

## How to start Administrator

```bash
docker-compose up
```
