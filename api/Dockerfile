FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app-api

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app-api .

EXPOSE 8000

CMD ["./app-api"]
