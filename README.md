# gtag

タグの追加・削除が面倒だったので作ってみた。

コマンドを実行する前にpullする仕様。

## コマンド

#### 最新のコミットでオートインクリメント
タグが存在しなかったらv0.0.1になる
```shell script
gtag -m i
```

#### 最新のタグを削除
```shell script
gtag -m d
```

#### タグを指定する
```shell script
gtag -m v -v v0.1.0
```
