FROM alpine:3.22.0

RUN apk update && \
    apk upgrade

RUN apk add neovim
