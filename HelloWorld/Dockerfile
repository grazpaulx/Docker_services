FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o helloworld helloworld.go

EXPOSE 8080

CMD ["./helloworld"]
