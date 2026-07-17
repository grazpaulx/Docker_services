FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o evenodd evenodd.go

EXPOSE 8080

CMD ["./evenodd"]
