FROM alpine:latest

COPY ./uramd /usr/bin/uramd

WORKDIR /root/.arkeo

RUN apk add --no-cache \
    bash \
    coreutils \
    jq \
    ca-certificates \
    curl

EXPOSE 1317
EXPOSE 26656
EXPOSE 26657

ENTRYPOINT [ "/bin/sh" ]

CMD [ "/usr/bin/uramd", "help" ]