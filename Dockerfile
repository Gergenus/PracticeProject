FROM golang:1.22

COPY ./ ./
RUN go build -o main .
CMD ["./cmd/main"]