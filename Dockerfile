FROM golang:1.22

WORKDIR /usr/local/src

COPY . .
RUN go mod download

RUN go build cmd/main.go

CMD ["./main"]