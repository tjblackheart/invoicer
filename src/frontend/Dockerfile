# prod build
FROM node:13-alpine as build
WORKDIR /app
ARG VUE_APP_ENDPOINT
ENV VUE_APP_ENDPOINT $VUE_APP_ENDPOINT
ENV VUE_APP_LOCALE de-DE
COPY src .
RUN npm install && npm run build

# prod srv
FROM alpine:edge as prod
WORKDIR /app
RUN apk add --update ca-certificates nginx && mkdir -p /run/nginx
COPY ./conf/nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/dist .
CMD ["nginx", "-g", "daemon off;"]
