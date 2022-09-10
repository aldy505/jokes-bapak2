import preprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-node';
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
    trailingSlash: 'never',
    files: {
      routes: './src/routes',
      assets: './static',
      hooks: {
        server: './src',
        client: './src'
      },
      lib: './src/lib',
    },
    adapter: adapter({
      out: "dist"
    }),
  },
};

export default config;
