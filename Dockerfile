FROM golang:1.20

WORKDIR /app

RUN go mod init render && go mod tidy

COPY . .

RUN go build -o main main.go

CMD ["./main"]
