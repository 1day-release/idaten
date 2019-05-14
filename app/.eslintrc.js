module.exports = {
  root: true,
  env: {
    node: true
  },
  'extends': [
    'plugin:vue/strongly-recommended',
    '@vue/standard'
  ],
  rules: {
    'no-console': 'warn',
    'no-debugger': 'warn'
  },
  parserOptions: {
    parser: 'babel-eslint'
  }
}
