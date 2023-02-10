FROM golang:alpine as builder
LABEL maintainer="Max Guzman <max.guzman@icloud.com>"
RUN apk --no-cache add alpine-sdk
RUN mkdir -p /build
ADD . /build
WORKDIR /build
RUN GOOS=linux GOARCH=amd64 go build -tags musl -o stickerfy ./cmd/.

FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/stickerfy /app/
ENTRYPOINT ["/app/stickerfy"]
