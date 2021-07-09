import preprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-static';
import { windi } from 'svelte-windicss-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: [windi(), preprocess({ postcss: false })],

	kit: {
		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',
		ssr: false,
		trailingSlash: 'never',
		adapter: adapter({
			// default options are shown
			pages: 'dist',
			assets: 'dist',
			fallback: '200.html'
		})
	}
};

export default config;
// Workaround until SvelteKit uses Vite 2.3.8 (and it's confirmed to fix the Tailwind JIT problem)
const mode = process.env.NODE_ENV;
const dev = mode === 'development';
process.env.TAILWIND_MODE = dev ? 'watch' : 'build';
