FROM golang:latest

WORKDIR /app
COPY $PWD/main.go /app
RUN go build main.go

EXPOSE 8000
ENTRYPOINT ["./main"]