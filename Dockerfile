FROM golang:1.23-alpine

WORKDIR /go-crud-api

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]