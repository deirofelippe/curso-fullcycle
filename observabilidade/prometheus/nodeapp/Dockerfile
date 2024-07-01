FROM node:20.15.0-alpine3.20

WORKDIR /app

COPY ./package.json ./package-lock.json ./

RUN npm ci

COPY ./ ./

USER node

CMD ["/usr/local/bin/npm","start"]
