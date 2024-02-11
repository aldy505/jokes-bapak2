# Jokes Bapak2 Client

The frontend.

## Development

```bash
# Install modules
$ npm install

# Run local server
$ npm run dev

# build everything
$ npm run build
```

> You can preview the built app with `npm run preview`, regardless of whether you installed an adapter. This should
_not_ be used to serve your app in production.

## Used packages

| Name                       | Version          | Type                 |
|----------------------------|------------------|----------------------|
| @sveltejs/kit              | `1.0.0-next.480` | Framework            |
| svelte                     | `3.50.1`         | Framework            |
| typescript                 | `4.8.3`          | Static type language |
| svelte-i18n                | `3.4.0`          | i18n Library         |
| svelte-windicss-preprocess | `4.2.8`          | CSS Library          |
| @fontsource/fira-mono      | `4.5.9`          | Webfont              |
| @fontsource/rubik          | `4.5.11`         | Webfont              |
| dotenv                     | `16.0.2`         | Utils                |
| @sentry/browser            | `7.12.1`         | Logging              |

## Directory structure

```
.
├── Dockerfile          - Docker image for client
├── package.json        - Meta information & dependencies
├── README.md           - You are here
├── src
│  ├── app.html         - HTML entry point
│  ├── components       - Svelte component files
│  ├── global.d.ts      - Global type definition for Typescript
│  ├── languages        - i18n localization database
│  ├── lib              - Logic & utilities
│  └── routes           - Svelte page files
├── static              - Static/public directory
├── svelte.config.js    - Svelte configuration file
├── tsconfig.json       - Typescript configuration file
├── windi.config.ts     - WindiCSS configuration file
└── package-lock.json   - Packages lock file
```

## `.env` configuration

```ini
VITE_NODE_ENV=development
VITE_API_ENDPOINT=
VITE_SENTRY_DSN=
```