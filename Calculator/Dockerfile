FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o calculator calculator.go

EXPOSE 8080

CMD ["./calculator"]
