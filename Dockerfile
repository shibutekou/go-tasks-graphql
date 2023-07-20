FROM golang:1.20
LABEL authors="vgekko"

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o todolist cmd/main.go

EXPOSE 8800
CMD ["./todolist"]