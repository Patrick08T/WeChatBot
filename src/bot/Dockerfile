# Compile stage
FROM golang:1.19.0 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /bot

# Final stage
FROM debian:buster

EXPOSE 8000

WORKDIR /
COPY --from=build-env /bot /

CMD ["/bot"]