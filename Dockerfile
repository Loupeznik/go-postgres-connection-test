FROM golang:1.20-alpine as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./app ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 8000

ENTRYPOINT ["./app"]
