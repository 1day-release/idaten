module.exports = {
  css: {
    loaderOptions: {
      sass: {
        data: `@import "@/styles/functions.scss";`
      }
    }
  },

  devServer: {
    overlay: {
      warnings: false,
      errors: true
    },
    proxy: {
      '/api': {
        target: 'https://api.idaten.1day-release.com',
        secure: true,
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
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
