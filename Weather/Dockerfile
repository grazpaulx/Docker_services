FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o weather weather.go

EXPOSE 8080

CMD ["./weather"]
