module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  extends: ['eslint:recommended', 'plugin:@typescript-eslint/recommended', 'prettier'],
  plugins: ['svelte3', '@typescript-eslint'],
  ignorePatterns: ['*.cjs'],
  overrides: [{ files: ['*.svelte'], processor: 'svelte3/svelte3' }],
  settings: {
    'svelte3/typescript': () => require('typescript'),
    'svelte3/ignore-warnings': ({ code }) => code === 'missing-declaration'
  },
  parserOptions: {
    sourceType: 'module',
    ecmaVersion: 2020
  },
  env: {
    browser: true,
    es2017: true,
    node: true
  },
  rules: {
    'no-undef': 'off',
    'array-callback-return': 'error',
    'no-constant-binary-expression': 'error',
    'no-self-compare': 'error',
    'no-template-curly-in-string': 'error',
    'no-unmodified-loop-condition': 'error',
    'no-unreachable-loop': 'error',
    'arrow-body-style': ['error', 'as-needed'],
    'block-scoped-var': 'error',
    curly: ['error', 'all'],
    'no-eval': 'error',
    'no-implied-eval': 'error',
    'no-var': 'error',
    'one-var': ['error', 'never'],
    'prefer-arrow-callback': 'error',
    'prefer-const': 'error',
    yoda: 'error',
    'array-bracket-newline': ['error', { multiline: true }],
    'brace-style': 'error',
    'no-shadow': 'error',
    'no-use-before-define': 'error',
    'dot-notation': 'error'
  }
};
