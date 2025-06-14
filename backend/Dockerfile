FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

RUN mkdir -p /app/logs && chmod 777 /app/logs

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"] 