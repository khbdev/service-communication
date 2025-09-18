FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Binary nomini aniq beramiz
RUN go build -o gateway main.go

# Container ichida shu binaryni ishlatamiz
CMD ["./gateway"]
