FROM golang:1.25

WORKDIR /app

COPY . .

RUN go build -o main ./src/main.go

EXPOSE 3001

CMD ["./main"]