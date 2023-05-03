FROM node:latest as builder

WORKDIR /app

COPY ./MUP/ .

EXPOSE 8004

RUN npm install

RUN npm run build

CMD ["node", "./out/index.js"]
