import preprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-static';
import { windi } from 'svelte-windicss-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // Consult https://github.com/sveltejs/svelte-preprocess
  // for more information about preprocessors
  preprocess: [
    windi({
      configPath: './windi.config.ts',
      preflights: false,
    }),
    preprocess({ postcss: false }),
  ],

  kit: {
    // hydrate the <div id="svelte"> element in src/app.html
    target: '#svelte',
    trailingSlash: 'never',
    files: {
      routes: './src/routes',
      assets: './static',
      hooks: './src',
      lib: './src/lib',
    },
    adapter: adapter({
      // default options are shown
      pages: 'dist',
      assets: 'dist',
    }),
  },
};

export default config;
