import { shallowMount } from '@vue/test-utils'
import Slide from '@/components/Slide.vue'

describe('基本表示', () => {
  it('markdownとページ数が入力された時、htmlが表示されるか', done => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        markdown: `# H1
        ## H2`,
        pageNumber: 2
      }
    })
    expect(wrapper.vm.markedHtml).toBeTruthy()
    wrapper.vm.$nextTick(() => {
      expect(wrapper.find('.page').text().length).toBeTruthy()
      done()
    })
  })
})

describe('サイズの途中変更', () => {
  it('比率固定モードのプロパティの変更を反映できるか', () => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        maxWidth: 500,
        maxHeight: 687
      }
    })
    wrapper.setProps({
      maxWidth: 1112,
      maxHeight: 687
    })
    expect(wrapper.vm.pageStyles['width']).toBe(1112)
  })
})
/*
describe('モードスタイルの適応', () => {
  it('比率可変モードスタイルが適応されるか', done => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        markdown: '# H1',
        width: 1112
      }
    })
    wrapper.vm.$nextTick(() => {
      expect(wrapper.find('.page').attributes().style.indexOf(';min-height:')).toBeGreaterThanOrEqual(0)
      done()
    })
  })

  it('比率固定モードスタイルが適応されるか', done => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        markdown: '# H1',
        maxWidth: 1112,
        maxHeight: 687
      }
    })
    wrapper.vm.$nextTick(() => {
      expect(wrapper.find('.page').attributes().style.indexOf(';height:')).toBeGreaterThanOrEqual(0)
      done()
    })
  })
})
*/

describe('比率可変モード(プロパティwidth指定)→font-size,min-heightが変化', () => {
  it('プロパティwidthが1112pxの場合、.pageのfont-sizeが18pxになるか', () => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        width: 1112
      }
    })
    expect(wrapper.vm.pageStyles['font-size']).toBe(18)
  })

  it('プロパティwidthが1112pxの場合、.pageのmin-heightが687pxになるか', () => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        width: 1112
      }
    })
    expect(wrapper.vm.pageStyles['min-height']).toBe(687)
  })

  it('プロパティwidthが500pxの場合、.pageのmin-heightが309px(500*0.618)になるか', () => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        width: 500
      }
    })
    expect(wrapper.vm.pageStyles['min-height']).toBe(309)
  })
})

describe('比率固定モード(プロパティmax-widthおよびmax-height指定の場合)→font-size,width,heightが変化', () => {
  it('プロパティmax-widthが1112px、max-heightが687pxの場合、.pageのwidthが1112px、heightが687pxになり、font-sizeが18pxになるか', () => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        maxWidth: 1112,
        maxHeight: 687
      }
    })
    expect(wrapper.vm.pageStyles['width']).toBe(1112)
    expect(wrapper.vm.pageStyles['height']).toBe(687)
    expect(wrapper.vm.pageStyles['font-size']).toBe(18)
  })

  it('プロパティmax-widthが500px、max-heightが687pxの場合、.pageのwidthが500px、heightが309px(500*0.618)になり、font-sizeが8px(500/61.8)になるか', () => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        maxWidth: 500,
        maxHeight: 687
      }
    })
    expect(wrapper.vm.pageStyles['width']).toBe(500)
    expect(wrapper.vm.pageStyles['height']).toBe(309)
    expect(wrapper.vm.pageStyles['font-size']).toBe(8)
  })

  it('プロパティmax-widthが1112px、max-heightが309pxの場合、.pageのwidthが500px(309/0.618)、heightが309pxになり、font-sizeが8px(500/61.8)になるか', () => {
    const wrapper = shallowMount(Slide, {
      propsData: {
        maxWidth: 1112,
        maxHeight: 309
      }
    })
    expect(wrapper.vm.pageStyles['width']).toBe(500)
    expect(wrapper.vm.pageStyles['height']).toBe(309)
    expect(wrapper.vm.pageStyles['font-size']).toBe(8)
  })
})
