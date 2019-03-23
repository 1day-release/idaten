module.exports = {  
  css: {
    loaderOptions: {
      sass: {
        data: `@import "@/styles/styles.scss";`
      }
    }
  },
  devServer: {
    overlay: {
      warnings: true,
      errors: true
    }
  },
  chainWebpack: config => {
    const svgRule = config.module.rule('svg')

    svgRule.uses.clear()

    svgRule
      .use('vue-svg-loader')
      .loader('vue-svg-loader')
  }
}
