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
      let slideType = 3
      if (markdown.indexOf('# ') === 0) {
        slideType = 1
      } else if (markdown.indexOf('## ') === 0 && markdown.split(/\n/).length === 1) {
        slideType = 2
      }

      let html = ''

      if (slideType === 1) {
        let title = ''
        let subtitle = ''
        let date = ''
        let to = ''
        let from = ''
        markdown.split(/\n/).forEach((line, lineNum) => {
          const cutStart = (line.indexOf(': ') >= 0) ? (line.indexOf(': ')) + 2 : 0
          if (lineNum === 0) {
            title = line.replace(/(^# )|(^#)/, '')
          } else if (lineNum === 1) {
            subtitle = line.slice(cutStart)
          } else if (lineNum === 2) {
            date = line.slice(cutStart)
          } else if (lineNum === 3) {
            to = line.slice(cutStart)
          } else if (lineNum === 4) {
            from = line.slice(cutStart)
          }
        })

        html = `<div class="slide slide-type1">
            <div class="slide-header">`

        if (to) {
          html += `<p class="to">${to}</p>`
        }

        html += `
        </div>
            <div class="slide-body">`

        if (title) {
          html += `<h1 class="cover-title">${title}</h1>`
        }

        if (subtitle) {
          html += `<p class="cover-sub-title">${subtitle}</p>`
        }

        html += `
            </div>
            <div class="slide-footer">`

        if (date) {
          html += `
              <p class="date">
                <time datetime="${date}">${date}</time>
              </p>`
        }

        if (from) {
          html += `<p class="from">${from}</p>`
        }

        html += `
            </div>
          </div>`
      } else if (slideType === 2) {
        let title = ''
        markdown.split(/\n/).forEach((line, lineNum) => {
          if (lineNum === 0) {
            title = line
          } else {
            html = line
          }
        })
        html = ((title) ? title + '\n' : '')
      } else {
        html = markdown
      }

      return html
    }
  }
}
