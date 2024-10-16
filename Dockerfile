FROM golang:1.20.5

MAINTAINER zazolla14<azola.zubizarreta@gmail.com>

WORKDIR /app
COPY . .
RUN go mod download

COPY ./pkg/config/config.sample.yaml ./pkg/config/config.yaml
RUN go build -o /app/vocagame-be-azola ./cmd/api

EXPOSE 8888
ENTRYPOINT ["/app/vocagame-be-azola"]
