FROM alpine:edge as go
WORKDIR /srv
COPY app .
RUN apk add --update go gcc g++ git ca-certificates &&\
    go build -o /srv/bin/invoicer cmd/web/*

##

FROM node:lts-alpine as node
WORKDIR /srv
COPY app/ui .
RUN apk add --update yarn && yarn && yarn build

##

FROM alpine
WORKDIR /srv
COPY --from=go /srv/bin/invoicer .
COPY --from=node /srv/dist ./ui/dist

VOLUME /srv/var

CMD ["./invoicer"]
