import firebase from 'firebase/app'
import firebaseApp from '@/firebaseApp'

const db = firebase.firestore(firebaseApp)

const defaultSlide = {
  markdown: `# サンプルについて
サブタイトル : サブタイトル
日付 : 2019-01-01
宛名 : 〇〇様
製作者 : 〇〇

## Slide1

### Slide2
1. Ordered List1
2. Ordered List2
`
}

const storageName = 'idaten.markdown'
export default class Slide {
  constructor () {
    const storage = localStorage.getItem(storageName)
    if (!storage) {
      localStorage.setItem(storageName, JSON.stringify(defaultSlide))
    }
  }

  async list () {
    return new Promise((resolve) => {
      firebase.auth().onAuthStateChanged((user) => {
        if (!user) return resolve([])
        db.collection('slides').where('uid', '==', user.uid).get().then((docs) => {
          const slides = []
          docs.forEach((doc) => {
            let slide = doc.data()
            slide.id = doc.id
            slides.push(slide)
          })
          resolve(slides)
        })
      })
    })
  }

  async get (id) {
    return new Promise((resolve) => {
      firebase.auth().onAuthStateChanged((user) => {
        if (!user) return resolve(JSON.parse(localStorage.getItem(storageName)))
        db.collection('slides').doc(id).get().then((doc) => {
          resolve(doc.data())
        })
      })
    })
  }

  async create (data) {
    if (!data) data = defaultSlide

    return new Promise((resolve) => {
      firebase.auth().onAuthStateChanged((user) => {
        const id = btoa(String.fromCharCode(...crypto.getRandomValues(new Uint8Array(64)))).substring(0, 64).replace(/\//g, '0') // Create random 64 characters text
        db.collection('slides').doc(id).set({
          uid: user.uid,
          markdown: data.markdown
        }).then((doc) => {
          resolve(id)
        })
      })
    })
  }

  async update (id, data) {
    return new Promise((resolve) => {
      firebase.auth().onAuthStateChanged((user) => {
        if (!user) {
          localStorage.setItem(storageName, JSON.stringify(data))
          return resolve()
        }
        db.collection('slides').doc(id).update({
          markdown: data.markdown
        }).then((doc) => {
          resolve()
        })
      })
    })
  }

  async delete (id) {
    return new Promise((resolve) => {
      firebase.auth().onAuthStateChanged((user) => {
        db.collection('slides').doc(id).delete().then((doc) => {
          resolve()
        })
      })
    })
  }
}
