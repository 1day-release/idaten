# API 仕様

## エンドポイント

```
https://api.idaten.1day-release.com/
```

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

#### 認証成功(200)
```
{
  "access_token": "xxxxxxxxxx"
}
```

#### 認証失敗 無効なコード(401)
```
{
  "errors": [
    {
      "message": "Invalid code."
    }
  ]
}
```

#### 認証失敗 ユーザー情報の取得失敗(401)
```
{
  "errors": [
    {
      "message": "Could not get user info."
    }
  ]
}
```

#### 認証失敗 アクセストークン更新失敗(401)
```
{
  "errors": [
    {
      "message": "Could not update access token."
    }
  ]
}
```

### 処理仕様
1. パラメータのcodeを元に、GoogleのOAuth認証を行う
   1. フォームデータを下記のデータ構造、ヘッダーのContent-Typeをapplication/x-www-form-urlencoded、メソッドをPOSTとして、https://accounts.google.com/o/oauth2/token にHTTP通信を行う
```
code: 取得したコード
client_id: Googleで取得したClientID
client_secret: Googleで取得したClientSecret
redirect_uri: idaten編集画面のURL
grant_type: authorization_code
```
5. OAuth認証に失敗した場合、ステータスコードを401、エラーメッセージを「Invalid code.」としてレスポンスを返す、処理終了
6. OAuth認証取得データのaccess_tokenを元に、ユーザー情報を取得する
   1. ヘッダーのContent-Typeをapplication/json、ヘッダーのAuthorizationを「Bearer {取得したaccess_token}」、メソッドをGETとして、https://www.googleapis.com/oauth2/v1/userinfo にHTTP通信を行う
7. Googleユーザー情報の取得に失敗した場合、ステータスコードを401、エラーメッセージを「Could not get user info.」としてレスポンスを返す、処理終了
8. ユーザー情報取得データのemailをemail、OAuth取得データのaccess_tokenをaccess_tokenとして、usersテーブルを更新する
9. 更新に失敗した場合、ステータスコードを401、エラーメッセージを「Could not update access token.」としてレスポンスを返す、処理終了
10. OAuth取得データのaccess_tokenをaccess_tokenとして、ステータスコードを200としてレスポンスを返す

## ログアウト

Google への認証情報を切る

### URL

[POST]

```
auth/logout
```

### Query Parameter

#### Header

|      パラメータ       |                      説明                         |
| :-----------------: | :----------------------------------------------: |
|    Authorization    |  「Bearer 」+ ログインAPIで発行されたアクセストークン   |

```
Authorization: Bearer xxxxxxxxxx
```

#### Body

### Response

#### ログアウト成功(200)
```
{}
```

#### ログアウト失敗 認証に失敗(401)
```
{
  "errors": [
    {
      "message": "Could not authenticate."
    }
  ]
}
```

#### ログアウト失敗 アクセストークン削除に失敗(401)
```
{
  "errors": [
    {
      "message": "Could not remove access token."
    }
  ]
}
```

### 処理仕様
1. 認証処理を行う(共通処理)
   1. ヘッダーのAuthorizationが存在しないまたは、Authorizationの先頭に「Bearer 」が存在しない場合、ステータスコードを401、エラーメッセージを「Could not authenticate.」としてレスポンスを返す、処理終了
   2. ヘッダーのAuthorizationの「Bearer 」を省いた文字列をaccess_tokenとして、usersテーブルを検索する
   3. ユーザーが存在しない場合、ステータスコードを401、エラーメッセージを「Could not authenticate.」としてレスポンスを返す、処理終了
2. ヘッダーのAuthorizationの「Bearer 」を省いた文字列をaccess_tokenとして、usersテーブルを検索する
3. ユーザーが存在しない場合、ステータスコードを401、エラーメッセージを「Could not remove access token.」としてレスポンスを返す、処理終了
4. 該当したユーザーのaccess_tokenを空として更新する
5. 更新に失敗した場合、ステータスコードを401、エラーメッセージを「Could not remove access token.」としてレスポンスを返す、処理終了
9. ステータスコードを200としてレスポンスを返す

## スライド一覧取得

### URL

[GET]

```
slides
```

### Query Parameter

#### Header

|      パラメータ       |                      説明                         |
| :-----------------: | :----------------------------------------------: |
|    Authorization    |  「Bearer 」+ ログインAPIで発行されたアクセストークン   |

```
Authorization: Bearer xxxxxxxxxx
```

#### Body

| パラメータ |                     説明                     |
| :--------: | :------------------------------------------: |
|   email    | メールアドレス（ex. 1day.release@gmail.com） |

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
  },
]
```

### 処理仕様
1. 認証処理を行う(共通処理)
2. slidesテーブルを一覧しゅとく　

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
