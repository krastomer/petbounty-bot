FROM golang:1.19 AS builder

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -a -installsuffix cgo -o main cmd/petbounty-api/main.go

FROM alpine:3.13

WORKDIR /app
COPY --from=builder /app/main .
RUN apk add dumb-init

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
EXPOSE 8080

CMD ["sh", "-c", "./main" ]