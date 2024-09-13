FROM golang:1.23-alpine

WORKDIR /go/src/app

COPY . .

RUN go build -o goflight main.go

CMD ["./goflight"]