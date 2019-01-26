# API 仕様

## ログイン

クライアント側から受け取ったトークンを Google の認証サーバへ転送する  
転送されて帰ってきた認証情報をクライアント側へ返却する

### URL

[GET]

```
api/auth/login
```

### Request

```
{
  code: "xxxxxxxxx"
}
```

### Response

```
{
}
```

## ログアウト

Google への認証情報を切る

### URL

[GET]

```
api/auth/logout
```

### Request

### Response

## スライド一覧

### URL

[GET]

```
api/slides-list
```

### Request

```
{
  "email": "1day.release@gmail.com"
}
```

### Response

```
[
  {
    "slide_id": "xxxxxxxxxxxxxxxx",
    "email": "1day.release@gmail.com",
    "share_mode": 0,
    "create_at": "2019-01-01 00:00:00",
    "update_at": "2019-01-02 00:00:00"
  },
  {
    "slide_id": "xxxxxxxxxxxxxxxx",
    "email": "1day.release@gmail.com",
    "share_mode": 0,
    "create_at": "2019-01-01 00:00:00",
    "update_at": "2019-01-02 00:00:00"
  },
]
```

## スライド作成

### URL

[POST]

```
api/slides-create
```

### Request

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
api/slides-get
```

### Request

```
{
  "email": "1day.release@gmail.com"
  "slide_id": "xxxxxxxxxxxxxxxx"
}
```

### Response

```
{
  "markdown": "# test",
  "share_mode": 0,
  "create_at": "2019-01-01 00:00:00",
  "update_at": "2019-01-02 00:00:00"
}
```

## スライド保存

### URL

[PUT]

```
api/slides-save
```

### Request

```
{
  "email": "1day.release@gmail.com"
  "slide_id": "xxxxxxxxxxxxxxxx"
  "markdown": "# test",
  "share_mode": 0
}
```

### Response

```
{
  "update_at": "2019-01-02 00:00:00"
}
```

## スライド削除

### URL

[DELETE]

```
api/slides-delete
```

### Request

```
{
  "email": "1day.release@gmail.com"
  "slide_id": "xxxxxxxxxxxxxxxx"
}
```

### Response

```
{
  "status": "200"
}
```
