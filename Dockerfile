FROM alpine:3.7

RUN apk -U add ca-certificates

EXPOSE 8080

ADD <APPLICATION_NAME>-linux-amd64 /usr/local/bin/<APPLICATION_NAME>

CMD ["<APPLICATION_NAME>"]
