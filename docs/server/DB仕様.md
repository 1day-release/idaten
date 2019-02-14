# DB 仕様

## users

|    フィールド名   |           用途              |
| :-------------: | :------------------------: |
|      email      |    Google のメールアドレス    |
|   access_token  |    認証用のアクセストークン    |
|     user_id     |      GoogleのユーザーID      |
|      name       |         フルネーム           |
|    given_name   |       ファーストネーム        |
|   family_name   |       ファミリーネーム        |
|     picture     |     アカウントのアイコン       |
|     locale      |          ロケール            |

## slides

| フィールド名 |              用途              |
| :----------: | :----------------------------: |
|   slide_id   | ランダムなスライド ID( 48 桁)  |
|    email     |    Google のメールアドレス     |
|   markdown   |  全てのマークダウンのテキスト  |
|  share_mode  | 0:非公開/1:公開閲覧/2:公開編集 |
|  create_at   |     スライドを作成した日時     |
|  update_at   |     スライドを更新した日時     |
