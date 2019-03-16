import request from 'sync-request'
import mixin from '@/mixin'
const getSlideHtml = mixin.methods.$_getSlideHtml

const testCasesURL = 'https://script.google.com/macros/s/AKfycbwTNzktMR8aTSvpXQFnxTqNHj93vs-2bzhESVHIVJg/dev'

const testCasesResponse = request('GET', testCasesURL)
const testCases = JSON.parse(testCasesResponse.getBody('utf8'))

describe('getSlideHtml', async () => {
  testCases.forEach(testCase => {
    it(testCase.title, () => {
      expect(getSlideHtml(testCase.input)).toBe(testCase.output)
    })
  })
})
