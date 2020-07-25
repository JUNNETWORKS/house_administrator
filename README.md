# ⚔️️ House Administrator ⚔️️

Administrator となりアンダーホームにいる Servant を操るのだ!

![](https://3.bp.blogspot.com/-HXQ1lQN8KtE/XCMQqRua9mI/AAAAAAABZ-w/Rm0nZLpU28MyRLW9__mNf3zkPy9IG2YDgCKgBGAs/s1600/Omake%2BGif%2BAnime%2B-%2BSword%2BArt%2BOnline%2B-%2BAlicization%2B-%2BEpisode%2B12%2B-%2BQuinella%2BPontifex%2BAdministrator.gif)

## 全体構成図

[Draw.io](https://drive.google.com/file/d/12RB3hiBf5S1zVZie_11tMnItpChCKPAT/view?usp=sharing)

API とフロントは分離する

## 使用技術

- GO (net/http, html/template, ...)

## 使い方

### Administrator の起動

```bash
go run administrator/main.go
```

## ルーティング

### バック

| Path                                     | Description                                                     | Methods          |
| ---------------------------------------- | --------------------------------------------------------------- | ---------------- |
| rooms                                    | 部屋一覧                                                        | GET, POST        |
| rooms/:room                              | 部屋のトップページ                                              | GET, DELETE, PUT |
| /rooms/:room/measurements                | 利用可能な計測値の種類を返す/登録する (温度や湿度,CO2 濃度など) | GET, POST        |
| /rooms/:room/measurements/temperature    | 部屋の温度を返す                                                | GET              |
| /rooms/:room/controllers                 | 利用可能な操作                                                  | GET, POST        |
| /rooms/:room/controllers/air-conditioner | エアコンの操作                                                  | POST             |

基本的にバックエンドサーバーではブラウザで必要なデータの取得や新たなデバイスの登録などを行う.

具体的に Servant からデータを Administrator に送る際には gRPC で送る

### フロント

| Path | Description  |
| ---- | ------------ |
| /    | トップページ |

何も決まっとらん

### grpc

これから

## 開発ロードマップ

1. DB 構築の設定
2. DB に対してデータを挿入, 参照する処理を実装
3. gRPC で servant からデータを受け取れるように gRPC のインターフェースの実装
4. 部屋の温度や湿度, エアコン制御などの回路の試作,開発
5. servant から Administrator にデータを送る処理を実装

## 参考にしたサイト
