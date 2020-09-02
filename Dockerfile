FROM golang:1.14

WORKDIR /
COPY . .

RUN go build -o main .

CMD ["./main"]