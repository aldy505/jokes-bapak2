FROM node:16.17.0-bullseye

RUN apt-get update && apt-get upgrade --no-install-recommends -y \
    && apt-get install -y --no-install-recommends curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN npm ci && npm run build && rm -rf node_modules && npm ci --omit=dev

EXPOSE ${PORT}

CMD ["node", "dist/index.js"]