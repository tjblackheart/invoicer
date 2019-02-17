FROM alpine:edge as go
WORKDIR /srv
COPY app .
RUN apk add --update go gcc g++ git ca-certificates &&\
    go build -o invoicer -ldflags "-X main.Version=0.1.0 -X main.Build=1" cmd/*

##

FROM node:lts-alpine as node
WORKDIR /srv
COPY app/ui .
RUN apk add --update yarn && yarn && yarn build

##

FROM alpine

WORKDIR /srv
COPY --from=go /srv/invoicer .
COPY --from=node /srv/dist ./ui/dist

VOLUME /srv/var

CMD ["./invoicer"]
