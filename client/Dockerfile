FROM node:16.17.0-bullseye

RUN apt-get update && apt-get upgrade --no-install-recommends -y \
    && apt-get install -y --no-install-recommends curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

ENV VITE_NODE_ENV="development"
ENV VITE_SERVER_API_ENDPOINT="https://jokesbapak2.reinaldyrafli.com/api/"
ENV VITE_BROWSER_API_ENDPOINT="http://api:5000/"
ENV VITE_SENTRY_DSN=""

RUN npm ci && npm run build && rm -rf node_modules && npm ci --omit=dev

EXPOSE ${PORT}

CMD ["node", "dist/index.js"]