module.exports = {
  root: true,
  env: {
    node: true,
  },
  'extends': [
    'plugin:vue/recommended',
    'eslint:recommended',
    '@vue/standard',
  ],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'semi': [ 'error', 'never' ],
    'comma-dangle': [ 'error', 'always-multiline' ],
    'vue/component-name-in-template-casing': [ 'error', 'kebab-case', { 'ignores': [] } ],
    'vue/html-closing-bracket-newline': [ 'error', { 'multiline': 'never' } ],
  },
  parserOptions: {
    parser: 'babel-eslint',
  },
}
