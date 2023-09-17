FROM alpine:latest
MAINTAINER Patrick <q506765973@gmail.com>

WORKDIR /go/bin

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ .
RUN CGO_ENABLED=0 GOARCH=64 GOOS=linux go build -a -o bot .

CMD ["/go/bin/bot"]