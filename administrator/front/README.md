# このディレクトリの役割

Administrator のフロントエンド

## 使用技術

### Main

- Next.js
  - TypeScript
  - react-query (API リクエストのキャッシュ用)
  - [Emotion](https://emotion.sh/docs/introduction): styled-component 的なやつ
- Tailwind CSS (CSS 書きたくないので)

### Test

- Jest (JS のテストフレームワーク)
- React Testing Library (React 公式推奨のテストライブラリ)

### Linter & Formatter

- ESLint
- Pretitter

### Design

- Figma (フロントエンドの Web デザインで使う)
- Storybook (UI コンポーネントのカタログが作れる) (UI コンポーネント作っても忘れそうなので導入したい) (余裕があればだけど)

## ディレクトリ構成

とりあえず初期はとくに何も考えず適当に作る. ある程度コード量が増えてきたら以下のように分割する.

presentational component と container component できっちり分離するようにする.

内部はには atomic design で構成する.

```
src
├── components  // Presentational Component
├── containers  // Container Component
└── pages
```

## ルーティング

| URL | 説明 |
| --- | ---- |
