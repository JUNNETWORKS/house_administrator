# このディレクトリの役割

Administrator のフロントエンド

## 使用技術

### Main

- Next.js
  - TypeScript
  - react-query (API リクエストのキャッシュ用)
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

presentational component と container component できっちり分離するようにする.

内部はには atomic design で構成する.

```
src
├── components  // Presentational Component
├── containers  // Container Component
└── pages
```
