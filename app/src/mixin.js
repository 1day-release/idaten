import authModel from './models/auth'
import slideModel from './models/slide'

export default {
  methods: {
    /**
     * スライドごとにMarkdownを区切る
     * @param markdown : マークダウン文字列
     * @return slideMarkdown : スライドの配列
     */
    $_getSlideMarkdown (markdown) {
      let slideMarkdown = ['']
      markdown.split(/\n/).forEach((value) => {
        let nowIndex = slideMarkdown.length - 1
        if (value.indexOf('# ') === 0 || value.indexOf('## ') === 0 || value.indexOf('### ') === 0) {
          slideMarkdown.push('')
          nowIndex++
        } else {
          slideMarkdown[nowIndex] += '\n'
        }
        slideMarkdown[nowIndex] += value
      })
      if (slideMarkdown[0] === '') slideMarkdown.shift()
      return slideMarkdown
    },

    /**
     * スライドのHTMLを取得する
     * @param markdown : マークダウン文字列
     * @return html : HTML
     */
    $_getSlideHtml (markdown) {
      return 'test'
    },

    /**
     * 認証APIを操作する
     * @return authModel
     */
    $_authModel: authModel,

    /**
     * スライドAPIを操作する
     * @return slideModel
     */
    $_slideModel: slideModel
  }
}
