module.exports = {
  transform: {
    '^.+\\.svelte$': [
      'svelte-jester',
      { preprocess: './svelte.config.test.cjs' }
    ],
    '^.+\\.ts$': 'ts-jest',
    '^.+\\.js$': 'ts-jest'
  },
  "roots": [
    "<rootDir>/src/lib/components", "<rootDir>/src/lib/stores"
  ],
  moduleFileExtensions: ['js', 'ts', 'svelte'],
  moduleNameMapper: {
    '^\\$lib(.*)$': '<rootDir>/src/lib$1',
    '^\\$app(.*)$': [
      '<rootDir>/.svelte-kit/dev/runtime/app$1',
      '<rootDir>/.svelte-kit/build/runtime/app$1'
    ]
  },
  setupFilesAfterEnv: ['<rootDir>/jest-setup.ts'],
  collectCoverageFrom: [
    '<rootDir>/src/**/**.{ts,svelte}',
    '!<rootDir>/src/lib/stores/products.ts',
    '!<rootDir>/src/lib/components/SearchBox.svelte',
    '!<rootDir>/src/lib/components/ProductsGrid.svelte'
  ],
  collectCoverage: true
}
