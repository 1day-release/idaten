import mixin from '@/mixin'
const methods = mixin.methods

describe('authModel', () => {
  const model = new methods.$_authModel()

  it('login', async () => {
    const result = await model.login()
    // expect(result).toBeTruthy()
  })
})

/*
describe('slideModel', () => {
  const model = new methods.$_slideModel()

  let slideId = null

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
})

describe('authModel2', () => {
  const model = new methods.$_authModel()

  it('logout', async () => {
    const result = await model.logout()
    expect(result).toBeTruthy()
  })
})
*/
