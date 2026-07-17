FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o currenttime currenttime.go

EXPOSE 8080

CMD ["./currenttime"]
