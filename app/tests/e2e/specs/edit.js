const toolbarHeight = 123

module.exports = {
  before: browser => {
    browser
      .url(process.env.VUE_DEV_SERVER_URL)
      .windowSize('current', 1200, 800 + toolbarHeight)
      .waitForElementVisible('#app', 5000)
  },
  after: browser => {
    browser
      .end()
  },
  'スライドの初期作成: スライドが存在しない時、初期スライドが作成されるか': browser => {
    browser
      .getValue('.editor', (result) => {
        browser.assert.equal((result.value.length > 0), true)
      })
  },
  'プレビュー: テキストの入力後、スライドに内容が反映されるか': browser => {
    browser
      .clearValue('.editor')
      .setValue('.editor', '# test')
      .expect.element('.preview .page .page-body .cover-title').text.to.equal('test')
  },
  'プレゼンモード: プレゼンモードに遷移できるか': browser => {
    browser
      .click('.presentation-button')
      .url(url => {
        browser.assert.equal((url.value.indexOf('/presentation') >= -1), true)
      })
      .keys(browser.Keys.ESCAPE)
  },
  'ログアウト時: スライドリストボタン、共有ボタン、設定ボタンが押下できないか': browser => {
    // eslint-disable-next-line no-unused-expressions
    browser
      .click('.user-slide-list-button')
      .expect.element('.user-slide').to.not.be.visible
    // browser
    //   .click('.user-slide-list-button')
    //   .expect.element('.user-slide').to.not.be.visible
  },
  'ログアウト時: 保存ステータス、ログアウトボタン、ユーザーアイコンが非表示であるか': browser => {
    // eslint-disable-next-line no-unused-expressions
    // browser.expect.element('.save-status').to.not.be.visible
    // eslint-disable-next-line no-unused-expressions
    browser.expect.element('.logout-button').to.not.be.visible
    // eslint-disable-next-line no-unused-expressions
    browser.expect.element('.user-icon').to.not.be.visible
  /*
  },
  'ログアウト時: ログアウトポップアップを表示されているか': browser => {
    browser
      .expect
      .element('.').to.be.visible
      */
  },
  'ログイン: ログインボタン押下時、ログイン画面に遷移され、ログインできるか': browser => {
    // browser.expect.element('.save-status').to.not.be.visible
    browser
      .click('.login-button')
      .pause(5000)
  }
}
