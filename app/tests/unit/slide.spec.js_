import request from 'sync-request'
import jsBeautify from 'js-beautify'
import mixin from '@/mixin'
const getSlideHtml = mixin.methods.$_getSlideHtml

const testCasesURL = 'https://script.google.com/macros/s/AKfycbxYk9NMXJbjxCdHMyTmoFH-mIVfoCgw2PXr_0NWIuFvp5dzfrE/exec'

const testCasesResponse = request('GET', testCasesURL)
const testCases = JSON.parse(testCasesResponse.getBody('utf8'))

function beautifyHtml (html) {
  const lines = html.split('\n')

  let str = ''
  lines.forEach(line => {
    str += line.replace(/^ +/g, '')
  })

  return jsBeautify.html(str)
}

describe('getSlideHtml', async () => {
  testCases.forEach(testCase => {
    if (!testCase.test_flg) return
    it(testCase.title, () => {
      const input = beautifyHtml(getSlideHtml(testCase.input))
      const output = beautifyHtml(testCase.output)

      expect(input).toBe(output)
    })
  })
})
