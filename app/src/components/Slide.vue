<template>
  <div class="slide" :class="['slide-type' + slideType]" v-html="marked(markdown)"></div>
</template>

<script>
import Marked from 'marked'

export default {
  name: 'Slide',
  data () {
    return {
    }
  },
  props: {
    markdown: String
  },
  created () {
  },
  methods: {
    marked (value) {
      let str = ''
      if (this.slideType === 1) {
        let title = ''
        let subtitle = ''
        let date = ''
        let to = ''
        let author = ''
        value.split(/\n/).forEach((line, lineNum) => {
          if (lineNum === 0) {
            title = line
          } else if (lineNum === 1) {
            const cutStart = line.indexOf(': ') || 0
            subtitle = line.slice(cutStart + 2)
          } else if (lineNum === 2) {
            const cutStart = line.indexOf(': ') || 0
            date = line.slice(cutStart + 2)
          } else if (lineNum === 3) {
            const cutStart = line.indexOf(': ') || 0
            to = line.slice(cutStart + 2)
          } else if (lineNum === 4) {
            const cutStart = line.indexOf(': ') || 0
            author = line.slice(cutStart + 2)
          }
        })
        str = ((to) ? to + '\n' : '') + ((title) ? title + '\n' : '') + ((subtitle) ? subtitle + '\n\n' : '') +
         ((date) ? date + '\n\n' : '') + ((author) ? author + '\n' : '') + '\n'
      } else if (this.slideType === 2) {
        let title = ''
        value.split(/\n/).forEach((line, lineNum) => {
          if (lineNum === 0) {
            title = line
          }
        })
        str = ((title) ? title + '\n' : '')
      } else {
        str = value
      }
      return Marked(str)
    }
  },
  computed: {
    slideType () {
      if (this.markdown.indexOf('# ') === 0) {
        return 1
      } else if (this.markdown.indexOf('## ') === 0) {
        return 2
      } else {
        return 3
      }
    }
  }
}
</script>

<style scoped lang="scss">
.slide {
  background: white;
  width: 100%;
  min-height: 70%;
  color: #000c36;
  &.slide-type1  /deep/ {
    padding: 5% 5%;
    box-sizing: border-box;
    p:nth-of-type(1) {
      font-size: 16px;
      margin-bottom: 25%;
    }
    h1 {
      font-size: 30px;
      margin-bottom: 2%;
    }
    p:nth-of-type(2) {
      margin-bottom: 25%;
    }
    p:nth-of-type(3) {
      color: #cbcdd6;
      margin-bottom: 1%;
    }
    p:nth-of-type(4) {
      color: #cbcdd6;
    }
  }
  &.slide-type2  /deep/ {
    h2 {
      margin-top: 30%;
      font-size: 30px;
      color: #000c36;
      text-align: center;
    }
  }
  &.slide-type3  /deep/ {
    padding: 10% 10%;
    box-sizing: border-box;
    h3 {
      margin-bottom: 5%;
      font-size: 30px;
      color: #000c36;
      text-align: center;
    }

    ol {
      list-style-type: decimal;
    }

    ul {
      list-style-type: disc;
    }
  }
}
</style>
