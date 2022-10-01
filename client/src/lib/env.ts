export default {
  SERVER_API_ENDPOINT: import.meta.env.VITE_SERVER_API_ENDPOINT || 'http://localhost:5000',
  BROWSER_API_ENDPOINT: import.meta.env.VITE_BROWSER_API_ENDPOINT || 'https://jokesbapak2.reinaldyrafli.com',
  SENTRY_DSN: import.meta.env.VITE_SENTRY_DSN || '',
  NODE_ENV: import.meta.env.VITE_NODE_ENV || 'development',
};
