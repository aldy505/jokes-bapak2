module.exports = {
	root: true,
	parser: '@typescript-eslint/parser',
	extends: ['eslint:recommended', 'plugin:@typescript-eslint/recommended', 'prettier',  'plugin:svelte/prettier'],
	plugins: ['@typescript-eslint'],
	ignorePatterns: ['*.cjs'],
	settings: {
		'svelte3/typescript': () => require('typescript')
	},
	overrides: [
		{
			files: ["*.svelte"],
			parser: "svelte-eslint-parser",
			// Parse the `<script>` in `.svelte` as TypeScript by adding the following configuration.
			parserOptions: {
				parser: "@typescript-eslint/parser",
			},
		},
	],
	parserOptions: {
		sourceType: 'module',
		ecmaVersion: 2019,
		extraFileExtensions: [".svelte"]
	},
	env: {
		browser: true,
		es2017: true,
		node: true
	}
};
