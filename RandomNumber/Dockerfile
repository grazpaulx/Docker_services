FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o random random.go

EXPOSE 8080

CMD ["./random"]
