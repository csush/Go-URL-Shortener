FROM golang:1.24.5 AS builder

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /shortener

FROM alpine:latest AS runner

COPY --from=builder /shortener /app/shortener

EXPOSE 5001

CMD ["/app/shortener"]