# idaten app 仕様

## 編集モード: /xxx/edit
- app
  - header
    - slide-list-show-button
    - brand-logo
    - presentation-button
    - user-icon
    - login-button
    - logout-button
  - user-slide-list
    - slide-list-hide-button
    - add-slide-button
    - user-slide
  - editor
    - popups
      - popup-logout
  - preview
    - slide
      - slide-page
      
## プレゼンモード: xxx/presentation
- presentation
  - slide
    - slide-page

### 仕様一覧
- ログイン時
  - appにlogoutクラスが削除される
    - app header login-button非表示
- 未ログイン時
  - appにlogoutクラスが付与される
    - app header user-icon非表示
    - app header logout-button非表示
    - app header slide-list-show-buttonが網かけ状態になる
    - app editor popups popup-logoutが表示
- app header slide-list-show-button押下
  - app slide-listにis_showクラスが付与される
    - app slide-listを表示する
- app header slide-list-hide-button押下
  - app slide-listのis_showクラスが削除される
- app header presentation-button押下
  - /xxx/presentationに別窓遷移する
- app header login-button押下
  - ログアウトAPIを実行する
- app user-slide-list user-slide押下
  - 選択されたスライドに同窓遷移する
- 現在のスライドのpreview slide slide-pageにis_nowクラスを付与する
  
  
## データ構造
```javascript
{
  markdown: ""
}
```

## 編集モード仕様
1. editorの変更のイベントを検知
2. editorの内容を取得する
3. データ構造を元に、markdownに内容を格納、json文字列に変換しidatenとしてローカルストレージに格納する

## プレゼンモード仕様
1. idatenをローカルストレージから取得、jsonパースを行い、markdownを取得する
2. markdownをmarked.jsを元にhtmlに変換
3. #または##または###で配列に分解
4. 配列を元に、スライドとして各セクションに分解し、内容を書き出す
5. 内容にスタイルを加える
