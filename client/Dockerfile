FROM node:16.13-bullseye

WORKDIR /app

COPY . .

RUN npm install && npm run build

EXPOSE ${PORT}

CMD ["npm", "run", "preview"]