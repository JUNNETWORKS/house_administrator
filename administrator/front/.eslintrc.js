module.exports = {
  env: {
    browser: true,
    es2020: true,
  },
  extends: [
    'plugin:react/recommended',
    'airbnb',
    'airbnb/hooks',
    'plugin:import/errors',
    'plugin:import/warnings',
    'plugin:import/typescript',
    'plugin:@typescript-eslint/eslint-recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:@typescript-eslint/recommended-requiring-type-checking',
    'plugin:prettier/recommended',
    'prettier',
    'prettier/@typescript-eslint',
    'prettier/react',
    'prettier/standard',
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 2020,
    project: './tsconfig.eslint.json',
    sourceType: 'module',
    tsconfigRootDir: __dirname,
  },
  plugins: [
    '@typescript-eslint',
    'import',
    'jsx-a11y',
    'prettier',
    'react-hooks',
  ],
  root: true,
  rules: {
    "no-use-before-define": "off",
    "@typescript-eslint/no-use-before-define": ["error"],
    'lines-between-class-members': [
      'error',
      'always',
      {
        exceptAfterSingleLine: true,
      },
    ],
    // Next.jsのLinkコンポーネント対応
    "jsx-a11y/anchor-is-valid": [ 
      "error", {
      "components": [ "Link" ],
      "specialLink": [ "hrefLeft", "hrefRight" ],
      "aspects": [ "invalidHref", "preferButton" ]
    }],
    // should be rewritten as `['error', { allowAsStatement: true }]` in ESLint 7 or later
    // SEE: https://github.com/typescript-eslint/typescript-eslint/issues/1184
    'no-void': 'off',
    'padding-line-between-statements': [
      'error',
      {
        blankLine: 'always',
        prev: '*',
        next: 'return',
      },
    ],
    '@typescript-eslint/no-unused-vars': [
      'error',
      {
        'vars': 'all',
        'args': 'after-used',
        'argsIgnorePattern': '_',
        'ignoreRestSiblings': false,
        'varsIgnorePattern': '_',
      },
    ],
    'import/extensions': [
      'error',
      'ignorePackages',
      {
        js: 'never',
        jsx: 'never',
        ts: 'never',
        tsx: 'never',
      },
    ],
    'react/jsx-filename-extension': [
      'error',
      {
        extensions: ['.jsx', '.tsx'],
      },
    ],
    "import/no-unresolved": [
      2,
      { ignore: ['\.(scss|less|css)$'] }
    ]
  },
  overrides: [
    {
      // Next.js のためにいくつかの項目を無効化する
      'files': ['*.js', '*.jsx', '*.ts', '*.tsx'],
      'rules': {
        'react/react-in-jsx-scope': 'off',
        'react/jsx-props-no-spreading': 'off',
      }
    },
    {
      'files': ['*.tsx'],
      'rules': {
        'react/prop-types': 'off',
      },
    },
  ],
  settings: {
    'import/resolver': {
      node: {
        extensions: [".js", ".jsx", ".ts", ".tsx", ".css", "json"],
        paths: ['src'],
      },
    },
    'import/ignore': {
      ignore: ['\.(scss|less|css)$']
    }
  },
};
