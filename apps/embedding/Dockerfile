FROM node:lts

RUN mkdir -p /opt/app
WORKDIR /opt/app

COPY package.json .
RUN mkdir -p src
COPY src/ ./src
RUN npm install
RUN npm run build
EXPOSE 3000
CMD [ "npm", "start"]