FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o server2 .
EXPOSE 8082
CMD ["/app/server2"]
