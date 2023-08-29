FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app/
COPY . /app/
RUN go build -o random-scheduler .

ENTRYPOINT ["/app//random-scheduler"]
