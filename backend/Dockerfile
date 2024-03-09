FROM golang:1.22

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.* ./
COPY .air.toml ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
