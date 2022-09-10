FROM node:16.17.0-bullseye

WORKDIR /app

COPY . .

RUN npm ci && npm run build && rm -rf node_modules && npm ci --omit=dev

EXPOSE ${PORT}

CMD ["node", "dist/index.js"]