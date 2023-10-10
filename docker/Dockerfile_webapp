FROM node:18-alpine as build
WORKDIR /app
COPY ./webapp .
RUN yarn
RUN yarn build

FROM node:18-alpine as deploy-node
WORKDIR /app
COPY --from=build /app/package.json /app/yarn.lock ./
COPY --from=build /app/target ./
RUN yarn --prod
ENV PORT=8080
CMD ["node", "index.js"]
