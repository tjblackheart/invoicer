# wkhtmltopdf
FROM aantonw/alpine-wkhtmltopdf-patched-qt:latest as wkhtmltopdf

# prod build
FROM golang:1.13-alpine as build
WORKDIR /go/src/app
COPY src .
RUN apk add --update gcc g++ && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/invoicer .

# prod out
FROM alpine:edge as prod
WORKDIR /app
RUN apk add --update ca-certificates wkhtmltopdf && mkdir -p /app/var/out
COPY --from=wkhtmltopdf /lib/libwkhtmltox.so.0.12.5 /usr/lib/wkhtmltox.so.0.12.5
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltoimage /usr/bin/
COPY --from=build /go/src/app/tpl/ ./tpl
COPY --from=build /go/src/app/bin/invoicer .
VOLUME [ "/app/var", "/app/var/out" ]
CMD ["./invoicer"]
