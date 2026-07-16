FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o greeting greeting.go

EXPOSE 8080

CMD ["./greeting"]
