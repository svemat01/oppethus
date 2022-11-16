FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o oppethus .

# ------------------------------

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/oppethus .

EXPOSE 8080

CMD ["./oppethus"]