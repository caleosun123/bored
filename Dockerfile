FROM golang:1.20

WORKDIR /app

# Initialize the module and get dependencies
RUN go mod init render && go mod tidy && go get -u github.com/go-sql-driver/mysql

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main main.go

# Run the Go application
CMD ["./main"]
