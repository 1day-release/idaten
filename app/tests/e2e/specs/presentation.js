const toolbarHeight = 123
const markdown = `
## 1
1
## 2
2
## 3
3
## 4
4`

module.exports = {
  before: browser => {
    browser
      .url(process.env.VUE_DEV_SERVER_URL)
      .windowSize('current', 1200, 800 + toolbarHeight)
      .waitForElementVisible('#app', 5000)
      .clearValue('.editor')
      .setValue('.editor', markdown)
      .click('.presentation-button')
  },
  after: browser => {
    browser
      .end()
  },
  'スライドサイズ: 読み込み時に、画面幅がwidth:1200pxの時に、スライドサイズが1100pxになるか': browser => {
    browser
      .getElementSize('.presentation-body .page', (size) => {
        browser.assert.equal(size.value.width, 1100)
      })
  },
  'スライドサイズ: 画面幅がwidth:1200pxの時に、画面幅をwidth:1000pxにリサイズした時に、スライドサイズが900pxになるか': browser => {
    browser
      .windowSize('current', 1000, 800 + toolbarHeight)
      .getElementSize('.presentation-body .page', (size) => {
        browser.assert.equal(size.value.width, 900)
      })
  },
  'ページ遷移: 右キー、Enter、クリックを押した時に、次のページに遷移するか': browser => {
    browser
      .keys(browser.Keys.RIGHT_ARROW)
      .assert.containsText('.presentation-body .page-now', '2')
      .keys(browser.Keys.ENTER)
      .assert.containsText('.presentation-body .page-now', '3')
      .click('.presentation-body')
      .assert.containsText('.presentation-body .page-now', '4')
  },
  'ページ遷移: 左キー、Delete、右クリックを押した時に、前のページに遷移するか': browser => {
    browser
      .keys(browser.Keys.LEFT_ARROW)
      .assert.containsText('.presentation-body .page-now', '3')
      .keys(browser.Keys.DELETE)
      .assert.containsText('.presentation-body .page-now', '2')
      .mouseButtonClick('right')
      .assert.containsText('.presentation-body .page-now', '1')
  },
  'ページ遷移: 最初のページの時、前のページに遷移できないか': browser => {
    browser
      .keys(browser.Keys.LEFT_ARROW)
      .assert.containsText('.presentation-body .page-now', '1')
  },
  'ページ遷移: 最後のページの時、次のページに遷移できないか': browser => {
    browser
      .keys(browser.Keys.RIGHT_ARROW) // 2
      .keys(browser.Keys.RIGHT_ARROW) // 3
      .keys(browser.Keys.RIGHT_ARROW) // 4
      .keys(browser.Keys.RIGHT_ARROW) // 4
      .assert.containsText('.presentation-body .page-now', '4')
  },
  'プレゼンテーション終了: エスケープ、戻るをクリックした時に、編集に戻るか': browser => {
    browser
      .keys(browser.Keys.ESCAPE)
      .url(url => {
        browser.assert.equal((url.value.indexOf('/presentation') === -1), true)
      })
      .click('.presentation-button')
      .click('.presentation-close-button')
      .url(url => {
        browser.assert.equal((url.value.indexOf('/presentation') === -1), true)
      })
  }
}
