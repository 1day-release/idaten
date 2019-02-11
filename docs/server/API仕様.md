# API 仕様

## エンドポイント

```
https://api.idaten.1day-release.com/
```

## ログイン

クライアント側から受け取ったトークンを Google の認証サーバへ転送する  
転送されて帰ってきた認証情報をクライアント側へ返却する

### URL

[GET]

```
auth/login
```

### Query Parameter

| パラメータ |             説明              |
| :--------: | :---------------------------: |
|    code    | Google から送信されたトークン |

```
?code=xxxxxxxxxxxx
```

### Response

```
{}
```

## ログアウト

Google への認証情報を切る

### URL

[GET]

```
auth/logout
```

### Query Parameter

### Response

## スライド一覧取得

### URL

[GET]

```
slides
```

### Query Parameter

| パラメータ |                     説明                     |
| :--------: | :------------------------------------------: |
|   email    | メールアドレス（ex. 1day.release@gmail.com） |

```
?email=1day.release@gmail.com
```

### Response

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
