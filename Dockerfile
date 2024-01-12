FROM golang:alpine3.19@sha256:2523a6f68a0f515fe251aad40b18545155135ca6a5b2e61da8254df9153e3648 as build

WORKDIR /usr/src/app

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src .
RUN go build -v -o /usr/local/bin/app .

FROM alpine:3.19.0@sha256:13b7e62e8df80264dbb747995705a986aa530415763a6c58f84a3ca8af9a5bcd as release

COPY --from=build /usr/local/bin/app /app

ENTRYPOINT ["/app"]

EXPOSE 8080
