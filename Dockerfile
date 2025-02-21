FROM golang:1.21.4-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -o web-server-task

FROM alpine:3.15

WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/web-server-task .
COPY --from=builder /app/config.json .
COPY migrations ./migrations
COPY --chmod=+x wait-for-postgres.sh .

EXPOSE 8080

CMD ["./wait-for-postgres.sh", "postgres:5432", "--", "./web-server-task"]

