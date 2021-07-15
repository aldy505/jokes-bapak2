# Jokes Bapak2 Client

Still work in progress

## Development

```bash
# Install modules
$ yarn install

# Run local server
$ yarn dev

# build everything
$ yarn build
```

> You can preview the built app with `yarn preview`, regardless of whether you installed an adapter. This should _not_ be used to serve your app in production.

## Used packages

| Name | Version | Type |
| --- | --- | --- |
| @sveltejs/kit | `1.0.0-next.129` | Framework |
| svelte | `3.38.3` | Framework |
| typescript | `4.3.5` | Static type language |
| svelte-i18n | `3.3.9` | i18n Library |
| svelte-windicss-preprocess | `4.0.12` | CSS Library |
| @fontsource/fira-mono | `4.5.0` | Webfont |
| @fontsource/rubik | `4.5.0` | Webfont |
| dotenv | `10.0.0` | Utils |
| @sentry/browser | `6.9.0` | Logging |

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
├── windi.config.js     - WindiCSS configuration file
└── yarn.lock           - Packages lock file
```

## `.env` configuration

```ini
VITE_NODE_ENV=development
VITE_API_ENDPOINT=
VITE_SENTRY_DSN=
```