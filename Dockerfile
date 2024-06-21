# Build stage
FROM golang:1.22.3-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server ./cmd/app/main.go

FROM alpine:latest

WORKDIR /go/src/app

COPY --from=builder /build/bin/server .
COPY --from=builder /build/.env .

CMD [ "./server" ]
