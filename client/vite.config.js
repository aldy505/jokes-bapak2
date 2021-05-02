import { defineConfig } from 'vite';
import { resolve } from 'path';
import eslint from '@rollup/plugin-eslint';

export default defineConfig({
  publicDir: 'public',
  root: './src',
  server: {
    port: 8000,
    cors: false,
  },
  plugins: [
    eslint({
      fix: true,
    }),
  ],
  build: {
    rollupOptions: {
      input: {
        main: resolve(__dirname, 'src/index.html'),
      },
    },
  },
});
