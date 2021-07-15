import * as Sentry from '@sentry/browser';
import env from './env';

Sentry.init({
  dsn: String(env.SENTRY_DSN),
  tracesSampleRate: 0.5,
});

export default Sentry;
