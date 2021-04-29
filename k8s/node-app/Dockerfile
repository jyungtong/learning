FROM node:alpine

ENV BASE_DIR=/app

WORKDIR ${BASE_DIR}

ADD ./package.json ${BASE_DIR}
ADD ./package-lock.json ${BASE_DIR}

RUN npm ci

ADD . ${BASE_DIR}

EXPOSE 8080

CMD node index.js
