FROM node:14.17.3-buster

WORKDIR /app

COPY . .

RUN npm install -g yarn
RUN yarn install
RUN yarn build
RUN yarn cache clean

EXPOSE ${PORT}

CMD ["yarn", "preview"]