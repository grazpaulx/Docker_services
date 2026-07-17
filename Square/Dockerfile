FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o square square.go

EXPOSE 8080

CMD ["./square"]
