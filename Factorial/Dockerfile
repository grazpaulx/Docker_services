FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o factorial factorial.go

EXPOSE 8080

CMD ["./factorial"]
