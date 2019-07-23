FROM alpine:edge as go
WORKDIR /srv
COPY app .
RUN apk add --update go gcc g++ git ca-certificates wget &&\
    go build -o /srv/bin/invoicer cmd/web/*

##

FROM node:lts-alpine as node
WORKDIR /srv
COPY app/ui .
RUN apk add --update yarn && yarn && yarn build

##

FROM aantonw/alpine-wkhtmltopdf-patched-qt:latest as wkhtmltopdf

##

FROM alpine:edge
WORKDIR /app
RUN apk add --update wkhtmltopdf

COPY --from=wkhtmltopdf /lib/libwkhtmltox.so.0.12.5 /usr/lib/wkhtmltox.so.0.12.5
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltoimage /usr/bin/

COPY --from=go /srv/tpl/ ./tpl/
COPY --from=go /srv/bin/invoicer .
COPY --from=node /srv/dist ./ui/dist

VOLUME /app/var

CMD ["./invoicer"]
