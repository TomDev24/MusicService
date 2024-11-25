FROM golang:1.22

COPY . .

RUN go build -o app ./cmd/music_lib/main.go
