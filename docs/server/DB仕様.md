# DB 仕様

## idaten-users

- Hash Key: user_id
- Range Key: email

| フィールド名 |           用途           |
| :----------: | :----------------------: |
|   user_id    |   Google のユーザー ID   |
|    email     | Google のメールアドレス  |
| access_token | 認証用のアクセストークン |
|     name     |        フルネーム        |
|  given_name  |     ファーストネーム     |
| family_name  |     ファミリーネーム     |
|   picture    |   アカウントのアイコン   |
|    locale    |         ロケール         |

## idaten-slides

- Hash Key: slide_id
- Range Key: email

| フィールド名 |              用途              |
| :----------: | :----------------------------: |
|   slide_id   | ランダムなスライド ID( 48 桁)  |
|    email     |    Google のメールアドレス     |
|   markdown   |  全てのマークダウンのテキスト  |
|  share_mode  | 0:非公開/1:公開閲覧/2:公開編集 |
|  create_at   |     スライドを作成した日時     |
|  update_at   |     スライドを更新した日時     |
