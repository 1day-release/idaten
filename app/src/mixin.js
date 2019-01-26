export default {
  methods: {
    /**
     * 文字列をケバブケース化する
     * @param str : 変換する文字列
     * @return str : ケバブケース化された文字列
     */
    $_getKebabCase (str) {
      return str.replace(/ /g, '-')
    },

    /**
     * 配列からランダムな要素を取得する
     * @param ary : 配列要素
     * @return item : ランダムに取得した要素
     */
    $_getArrayRandom (ary) {
      return ary[Math.floor(Math.random() * ary.length)]
    },

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
    }
  }
}
