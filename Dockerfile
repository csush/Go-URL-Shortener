FROM golang:1.24.5

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /shortener

EXPOSE 5001

CMD ["/shortener"]