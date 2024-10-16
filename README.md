# VocaGame-BackendTest-Azola

An example of gin contains many useful features for e-commerce websites

## How to run

### Required Environment

- Go 1.20
- Postgres
- Redis

### Config

- Copy config file: `cp pkg/config/config.sample.yaml pkg/config/config.yaml`
- You should modify `pkg/config/config.yaml`

```yaml
environment: production
http_port: 8888
auth_secret: auth_secret
database_uri: postgres://postgres:1234@postgres:5432/postgres
redis_uri: localhost:6379
redis_password:
redis_db: 0
```

### Run

```shell script
$ go run cmd/api/main.go
```

```docker compose
$ docker-compose up
```

```
2023-09-12T15:18:36.684+0700    INFO    http/server.go:58       HTTP server is listening on PORT: 8888
```

### Test

```shell script
$ go test
```

### Test with Coverage

```shell script
go test -timeout 9000s -a -v -coverprofile=coverage.out -coverpkg=./... ./...
```

**or**

```shell script
make unittest
```

### Document

- API document at: `http://localhost:8888/swagger/index.html`

### Tech stack

- Restful API
- DDD
- Gorm
- Swagger
- Logging
- Jwt-Go
- Gin-gonic
- Redis
- Postgres

### What's next?

- gRPC functions for products and orders
- Push message to notify place order successfully
- Payment with PayPal
- Define error response wrapper
