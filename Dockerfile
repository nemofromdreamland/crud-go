FROM golang:1.24 AS builder

WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
  go build -o crud-go main.go init_dependencies.go

FROM golang:1.24-alpine AS runner

RUN adduser -D nemo
RUN mkdir /app

COPY --from=builder /app/crud-go /app/crud-go

RUN chown -R nemo:nemo /app
RUN chmod +x /app/crud-go

WORKDIR /app

EXPOSE 8080

USER nemo

CMD ["./crud-go"]
