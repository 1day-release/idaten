import moment from 'moment'

const slidesStorageName = 'idaten.slides'
export default class Slide {
  constructor () {
    const slides = this.list()
    if (!slides) {
      localStorage.setItem(slidesStorageName, JSON.stringify([]))
      this.create({
        cover: '# サンプルについて',
        markdown: `# サンプルについて
サブタイトル : サブタイトル
日付 : 2019-01-01
宛名 : 〇〇様
製作者 : 〇〇

## Slide1

### Slide2
1. Ordered List1
2. Ordered List2
`,
        share_mode: 0
      })
    }
  }

  list () {
    const slides = JSON.parse(localStorage.getItem(slidesStorageName))
    return slides
  }

  get (id) {
    const slides = this.list()

    let result = null
    slides.forEach(slide => {
      if (slide.slide_id === id) result = slide
    })
    return result
  }

  create (data) {
    let slides = this.list()

    const id = btoa(String.fromCharCode(...crypto.getRandomValues(new Uint8Array(64)))).substring(0, 64).replace(/\//g, '0') // Create random 64 characters text
    slides.push({
      slide_id: id,
      cover: data.cover,
      markdown: data.markdown,
      share_mode: data.share_mode,
      create_at: moment().format('YYYY-MM-DD HH:mm:ss'),
      updated_at: moment().format('YYYY-MM-DD HH:mm:ss')
    })
    localStorage.setItem(slidesStorageName, JSON.stringify(slides))

    return {
      slide_id: id
    }
  }

  update (id, data) {
    let slideData = this.get(id)

    if (data.cover) slideData.cover = data.cover
    if (data.markdown) slideData.markdown = data.markdown
    if (data.share_mode) slideData.share_mode = data.share_mode
    slideData.updated_at = moment().format('YYYY-MM-DD HH:mm:ss')

    let slides = this.list()
    slides.forEach((slide, slideIndex) => {
      if (slide.slide_id === id) slides[slideIndex] = slideData
    })
    localStorage.setItem(slidesStorageName, JSON.stringify(slides))

    return {
      slide_id: id
    }
  }

  delete (id) {
    let slides = this.list()
    slides.forEach((slide, slideIndex) => {
      if (slide.slide_id === id) slide.splice(slideIndex, 1)
    })
    localStorage.setItem(slidesStorageName, JSON.stringify(slides))

    return {
      slide_id: id
    }
  }
}
