FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o server1 .
EXPOSE 8081
CMD ["/app/server1"]
