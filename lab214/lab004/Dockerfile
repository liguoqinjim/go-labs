FROM golang:latest

#ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY $PWD/main.go /app
COPY $PWD/go.mod /app
RUN go build main.go

ENTRYPOINT ["./main"]