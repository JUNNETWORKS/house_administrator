# ⚔️️ House Administrator ⚔️️

Administrator となりアンダーホームにいる Servant を操るのだ!

![](https://3.bp.blogspot.com/-HXQ1lQN8KtE/XCMQqRua9mI/AAAAAAABZ-w/Rm0nZLpU28MyRLW9__mNf3zkPy9IG2YDgCKgBGAs/s1600/Omake%2BGif%2BAnime%2B-%2BSword%2BArt%2BOnline%2B-%2BAlicization%2B-%2BEpisode%2B12%2B-%2BQuinella%2BPontifex%2BAdministrator.gif)

## 全体構成図

[Draw.io](https://drive.google.com/file/d/12RB3hiBf5S1zVZie_11tMnItpChCKPAT/view?usp=sharing)

API とフロントは分離する

## 使用技術

- [gin](https://github.com/gin-gonic/gin): Go の Web フレームワーク
- [Vue.js](https://github.com/gin-gonic/gin): フロント担当

## 使い方

### Administrator の起動

```bash
go run administrator/main.go
```

## ルーティング

### フロント

- `/`: トップページ

### バック

gRPC で実装予定.

## 開発ロードマップ

1. DB 構築の設定
2. DB に対してデータを挿入, 参照する処理を実装
3. gRPC で servant からデータを受け取れるように gRPC のインターフェースの実装
4. 部屋の温度や湿度, エアコン制御などの回路の試作,開発
5. servant から Administrator にデータを送る処理を実装

## 参考にしたサイト

- [gin を最速でマスターしよう](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)
- [gin GitHub](https://github.com/gin-gonic/gin)
