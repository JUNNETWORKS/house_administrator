# House Administrator

[Draw.io](https://drive.google.com/file/d/12RB3hiBf5S1zVZie_11tMnItpChCKPAT/view?usp=sharing)

API とフロントは分離せず, MVC アーキテクチャで構築する.

## 使用ライブラリ

- [gin](https://github.com/gin-gonic/gin): Go の Web フレームワーク
- [Vue.js](https://github.com/gin-gonic/gin): フロント担当

## 使い方

```bash
go run src/main.go
```

## ルーティング

### フロント

- `/`: トップページ

### バック

gRPC で実装予定.

## 参考にしたサイト

- [gin を最速でマスターしよう](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)
- [gin GitHub](https://github.com/gin-gonic/gin)
