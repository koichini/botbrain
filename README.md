# botbrain

## TODO

[x] JSON ファイルの更新 - ID を PK として挙動制御(ポインタメソッドを使用して)

[x] firestore の接続とその関数化エラー処理

[ ] エラー処理

[ ] Test ファイル作成

[ ] GCP へデプロイ

## Running

```
go run .
```

## Based reference page

https://go.dev/doc/tutorial/web-service-gin

## Firestore credential Document

https://cloud.google.com/docs/authentication/provide-credentials-adc?hl=ja#local-dev

## MEMO

- 初期で扱ったテストデータは`db/init.json`(不要なタイミングで削除)
