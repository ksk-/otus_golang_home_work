FROM alpine:3.9

RUN apk add curl
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

ENTRYPOINT [ "/usr/local/bin/goose" ]
