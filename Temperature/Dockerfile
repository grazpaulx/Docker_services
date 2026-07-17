FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o temperature temperature.go

EXPOSE 8080

CMD ["./temperature"]
