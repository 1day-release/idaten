# idaten app 仕様

## 編集モード: /xxx/edit
- app
  - header
    - slide-list-display-button
    - brand-logo
    - presentation-button
    - user-icon
    - login-button
    - logout-button
  - slide-list
    - slide
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
    - app header slide-list-display-buttonが網かけ状態になる
    - app editor popups popup-logoutが表示
- app header slide-list-display-button押下
  - app slide-listにis-showクラスが付与される
    - app slide-listを表示する
  - 再度押下すると非表示

- app header presentation-button押下
  - /xxx/presentationに別窓遷移する
  
- app header login-button押下
  - ログアウトAPIを実行する
  
- app slide-list slide押下
  - 選択されたスライドに遷移する
  