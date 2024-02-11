import { defineConfig } from 'astro/config';
import UnoCSS from '@unocss/astro'

// https://astro.build/config
export default defineConfig({
  i18n: {
    defaultLocale: "en",
    locales: ["en", "id"]
  },
  integrations: [
    UnoCSS({
      injectReset: true
    }),
  ],
});
