FROM alpine
RUN apk update && apk add --update --no-cache ca-certificates bash
COPY mmesh-node /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/mmesh-node"]
