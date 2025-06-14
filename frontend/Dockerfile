FROM node:20-alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

EXPOSE 5173

#JUST FOR TESTING
CMD ["npm", "run", "dev"]