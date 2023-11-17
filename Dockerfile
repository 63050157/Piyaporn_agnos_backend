FROM golang:1.21.4-alpine3.17
WORKDIR /app
COPY . .
RUN go build -o api
EXPOSE 8080
CMD ["./api"]
