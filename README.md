## 概要

- レイヤードアーキテクチャ+DDD を導入した API
- 思想は以下 qiita 参照
  https://qiita.com/karamaru/items/74880b29a054bdeb356c

## エンドポイント

### POST /signup

- ユーザー作成処理
- リクエストボディの Name から作成し、自動生成されたトークンを response
- ユーザー名は 2 文字以上 10 文字以下

### GET /account

- ユーザー取得処理
- Header「x-token」のトークンに応じたユーザーの ID と Name を response

### PATCH /account

- ユーザー更新処理
- Header「x-token」のトークンに応じたユーザーの name をリクエストボディの Name で更新
- ユーザー名は 2 文字以上 10 文字以下
