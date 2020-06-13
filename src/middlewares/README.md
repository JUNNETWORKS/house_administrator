# このディレクトリの役割

gin のリクエストに対する処理の流れは以下のようなものである.

```plain
Request -> Route Parser -> Middleware -> Route Handler -> Middleware -> Response
```

実際処理する関数の前に必ず Middleware は通る形になる.

なので, アクセスログの記録などに使える.
