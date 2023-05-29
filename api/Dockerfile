FROM golang:alpine as builder
LABEL maintainer="Max Guzman <max.guzman@icloud.com>"
RUN mkdir -p /build
ADD ./api /build
WORKDIR /build
RUN apk add --no-cache git gcc musl-dev
RUN GOOS=linux GOARCH=amd64 go build -tags musl -o stickerfy .

FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/stickerfy /app/
ENTRYPOINT ["/app/stickerfy"]
