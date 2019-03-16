# API 仕様

## エンドポイント

```
https://api.idaten.1day-release.com/
```

## ステータスコード

### 成功

| HTTP Code | 説明     |
| --------- | -------- |
| 200       | 成功     |
| 201       | 登録成功 |

### 失敗

レスポンスのボディに code と message でエラーの原因を配列で含める  
複数存在する場合は配列内部に複数記述する  
code は「分類アルファベット + 発番」により構成される

```
{
  "errors": [
    {
      "code": "a1",
      "message": "Invalid code."
    },
    {
      "code": "a2",
      "message": "Could not get user info."
    },
    {
      "code": "a3",
      "message": "Could not update access token."
    }
  ]
}
```

#### 認証(a)

| Error Code | HTTP Code | 説明                           |
| ---------- | --------- | ------------------------------ |
| a1         | 403       | Invalid code.                  |
| a2         | 403       | Could not get user info.       |
| a3         | 403       | Could not update access token. |
| a4         | 403       | Could not authenticate.        |
| a5         | 403       | Could not remove access token. |

#### 取得(b)

| Error Code | HTTP Code | 説明 |
| ---------- | --------- | ---- |
| b1         | 404       |      |

#### 登録(c)

| Error Code | HTTP Code | 説明 |
| ---------- | --------- | ---- |
| c1         | 400       |      |

#### 削除(d)

| Error Code | HTTP Code | 説明 |
| ---------- | --------- | ---- |
| d1         | 400       |      |
| d1         | 404       |      |

## ログイン

クライアント側から受け取ったトークンを Google の認証サーバへ転送する
転送されて帰ってきた認証情報をクライアント側へ返却する

### URL

[POST]

```
auth/login
```

### Query Parameter

#### Header

#### Body

| パラメータ |             説明              |
| :--------: | :---------------------------: |
|    code    | Google から送信されたトークン |

```
{
  "code": "xxxxxxxxxxxx"
}
```

### Response

```
{
  "access_token": "xxxxxxxxxx"
}
```

### 処理仕様

1. パラメータの code を元に、Google の OAuth 認証を行う
   1. フォームデータを下記のデータ構造、ヘッダーの Content-Type を application/x-www-form-urlencoded、メソッドを POST として、https://accounts.google.com/o/oauth2/token に HTTP 通信を行う

```
code: 取得したコード
client_id: Googleで取得したClientID
client_secret: Googleで取得したClientSecret
redirect_uri: idaten編集画面のURL
grant_type: authorization_code
```

2. OAuth 認証に失敗した場合、Error Code を a1 としてレスポンスを返す、処理終了
3. OAuth 認証取得データの access_token を元に、ユーザー情報を取得する
   1. ヘッダーの Content-Type を application/json、ヘッダーの Authorization を「Bearer {取得した access_token}」、メソッドを GET として、https://www.googleapis.com/oauth2/v1/userinfo に HTTP 通信を行う
4. Google ユーザー情報の取得に失敗した場合、Error Code を a2 としてレスポンスを返す、処理終了
5. OAuth 取得データの access_token を access_token、ユーザー情報取得データを各カラムに分解して、users テーブルを更新する
6. 更新に失敗した場合、Error Code を a3 としてレスポンスを返す、処理終了
7. OAuth 取得データの access_token を access_token として、Http Code を 200 としてレスポンスを返す

## ログアウト

Google への認証情報を切る

### URL

[POST]

```
auth/logout
```

### Query Parameter

#### Header

| パラメータ    | 説明                                                   |
| ------------- | ------------------------------------------------------ |
| Authorization | 「Bearer 」+ ログイン API で発行されたアクセストークン |

```
Authorization: Bearer xxxxxxxxxx
```

#### Body

### Response

```
{
  "status": "OK"
}
```

### 処理仕様

1. 認証処理を行い、ユーザー構造体を取得する(共通処理)
   1. ヘッダーの Authorization が存在しないまたは、Authorization の先頭に「Bearer 」が存在しない場合、Error Code を a4 としてレスポンスを返す、処理終了
   2. ヘッダーの Authorization の「Bearer 」を省いた文字列を access_token として、users テーブルを検索する
   3. ユーザーが存在しない場合、Error Code を a4 としてレスポンスを返す、処理終了
   4. ヘッダーの Authorization の「Bearer 」を省いた文字列を access_token として、users テーブルを検索する
   5. ユーザーが存在しない場合、Error Code を a4 としてレスポンスを返す、処理終了
   6. ユーザー構造体を返す
2. ユーザー構造体の access_token を空にして、users テーブルを更新する
3. 更新に失敗した場合、Error Code を a5 としてレスポンスを返す、処理終了
4. Http Code を 200 としてレスポンスを返す

## スライド一覧取得

### URL

[GET]

```
slides
```

### Query Parameter

#### Header

| パラメータ    | 説明                                                   |
| ------------- | ------------------------------------------------------ |
| Authorization | 「Bearer 」+ ログイン API で発行されたアクセストークン |

```
Authorization: Bearer xxxxxxxxxx
```

#### Body

| パラメータ | 説明                                        |
| ---------- | ------------------------------------------- |
| email      | メールアドレス（ex.1day.release@gmail.com） |

```
?email=1day.release@gmail.com
```

### Response

#### 取得成功(200)

```
[
  {
    "slide_id": "xxxxxxxxxxxxxxxx",
    "email": "1day.release@gmail.com",
    "share_mode": 0,
    "created_at": "2019-01-01 00:00:00",
    "updated_at": "2019-01-02 00:00:00"
  },
  {
    "slide_id": "xxxxxxxxxxxxxxxx",
    "email": "1day.release@gmail.com",
    "share_mode": 0,
    "created_at": "2019-01-01 00:00:00",
    "updated_at": "2019-01-02 00:00:00"
  }
]
```

### 処理仕様

1. 認証処理を行い、ユーザー構造体を取得する(共通処理)
2. ユーザー構造体を元に、email で slides テーブルを絞り込む
3. 取得に失敗した場合、レスポンスデータを空配列、Http Code を 200 としてレスポンスを返す、処理終了
4. レスポンスデータを取得したデータ、Http Code を 200 としてレスポンスを返す

## スライド作成

### URL

[POST]

```
slides
```

### Request Body

| パラメータ |                     説明                     |
| :--------: | :------------------------------------------: |
|   email    | メールアドレス（ex. 1day.release@gmail.com） |
| share_mode |       0:非公開，1:公開閲覧，2:公開編集       |

```
{
    "email": "1day.release@gmail.com",
    "share_mode": 0
}
```

### Response

```
{
    "slide_id": "xxxxxxxxxxxxxxxx"
}
```

## スライド取得

### URL

[GET]

```
slides/{slide_id}
```

### Query Parameter

| パラメータ | 説明 |
| :--------: | :--: |
|    なし    |      |

```
URLのみ
```

### Response

```
[
  {
    "slide_id": "xxxxxxxxxxxxxxxx"
    "markdown": "# test",
    "share_mode": 0,
    "created_at": "2019-01-01 00:00:00",
    "updated_at": "2019-01-02 00:00:00"
  }
]
```

## スライド保存

### URL

[PUT]

```
slides/{slide_id}
```

### Request Body

| パラメータ |                     説明                     |
| :--------: | :------------------------------------------: |
|   email    | メールアドレス（ex. 1day.release@gmail.com） |
|  markdown  |            マークダウンのテキスト            |
| share_mode |       0:非公開，1:公開閲覧，2:公開編集       |

```
{
    "email": "1day.release@gmail.com"
    "markdown": "# test",
    "share_mode": 0
}
```

### Response

```
{
    "updated_at": "2019-01-02 00:00:00"
}
```

## スライド削除

### URL

[DELETE]

```
slides/{slide_id}
```

### Request Body

| パラメータ |                     説明                     |
| :--------: | :------------------------------------------: |
|   email    | メールアドレス（ex. 1day.release@gmail.com） |

```
{
    "email": "1day.release@gmail.com"
}
```

### Response

```
{
    "status": "200"
}
```
