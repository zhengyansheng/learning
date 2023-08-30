FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

# Fetch dependencies.
WORKDIR /app/
COPY . /app/

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o random-scheduler .

# Run the Go Gin binary.
ENTRYPOINT ["/app//random-scheduler"]
