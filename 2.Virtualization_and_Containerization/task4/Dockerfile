FROM ubuntu:20.04
COPY ./go ./go
WORKDIR /app
COPY ./webapp/* ./
RUN /go/bin/go build -o webapp main.go
CMD ["./webapp"]
