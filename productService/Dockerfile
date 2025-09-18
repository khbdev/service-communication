FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o productService main.go

EXPOSE 8083

CMD ["./productService"]
