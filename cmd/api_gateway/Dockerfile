FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o app cmd/api_gateway/*.go

FROM alpine
COPY --from=builder /app /app
WORKDIR /app
EXPOSE 8080
CMD ["/app/app"]