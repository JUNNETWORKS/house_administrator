# 各ディレクトリの役割

- `controllers/`: `src/main.go` から来た HTTP リクエストの具体的な処理を実装している.
- `services/`: `routes`から要求された DB 操作などを行い,結果を返す
- `models/`: DB 構造に合わせた構造体(structure)の定義などを行う.
- `middlewares/`: ミドルウェア
- `views/`: Vue.js を使ったフロントエンド部分
