# wkhtmltopdf
FROM aantonw/alpine-wkhtmltopdf-patched-qt:latest as wkhtmltopdf

#dev
FROM golang:1.13-alpine as dev
WORKDIR /go/src/app
COPY src .
COPY ./conf/runner.conf /runner.conf
RUN apk add --update gcc g++ wkhtmltopdf git && \
    go get github.com/pilu/fresh && \
    chown -R 1000.1000 /go
COPY --from=wkhtmltopdf /lib/libwkhtmltox.so.0.12.5 /usr/lib/wkhtmltox.so.0.12.5
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltoimage /usr/bin/
CMD [ "fresh", "-c", "/runner.conf" ]

# prod
# TODO
