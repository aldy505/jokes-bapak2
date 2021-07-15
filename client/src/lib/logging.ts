import * as Sentry from '@sentry/browser';
import env from './env';

Sentry.init({
  dsn: String(env.SENTRY_DSN) || '',
  enabled: String(env.NODE_ENV) === 'production',
  tracesSampleRate: 0.5,
});

export default Sentry;
