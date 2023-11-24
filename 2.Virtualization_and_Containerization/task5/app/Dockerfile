FROM golang:1.21.2
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o app ./cmd/main
CMD ["./app"]
