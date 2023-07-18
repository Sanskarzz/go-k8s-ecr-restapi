FROM golang:1.19.0

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download

RUN go mod tidy

CMD ["air", "./cmd/main.go", "-b", "0.0.0.0"]