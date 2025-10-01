FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o server2 .
CMD ["./server2"]
