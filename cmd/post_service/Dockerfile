FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o app cmd/post_service/*.go

FROM alpine
COPY --from=builder /app /app
EXPOSE 9001
WORKDIR /app
CMD ["/app/app"]