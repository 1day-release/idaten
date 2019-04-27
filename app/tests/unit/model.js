
/*
import AuthModel from '@/models/auth'

describe('authModel', () => {
  const auth = new AuthModel()

  it('login', async () => {
    const result = await auth.login()
    expect(result).toBeTruthy()
  })
})
*/

import SlideModel from '@/models/slide'

describe('slideModel', () => {
  const model = new SlideModel()

  it('constructor', () => {
  })

  /*
  it('create', async () => {
    const result = await model.create({})
    expect(result).toBeTruthy()
    slideId = result.slide_id
  })

  it('update', async () => {
    const result = await model.update(slideId, {})
    expect(result).toBeTruthy()
  })

  it('list', async () => {
    const result = await model.list()
    expect(result.length).toBeTruthy()
  })

  it('get', async () => {
    const result = await model.get(slideId)
    expect(result).toBeTruthy()
  })

  it('delete', async () => {
    const result = await model.delete(slideId)
    expect(result).toBeTruthy()
  })

  it('get', async () => {
    const result = await model.get(slideId)
    expect(result).toBeFalsy()
  })
  */
})
