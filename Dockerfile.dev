FROM golang:1.25

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .air.docker.toml .air.toml

EXPOSE 3001

CMD ["air"]