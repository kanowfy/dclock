FROM golang:1.19.1-alpine3.15

WORKDIR /app
COPY . . 

RUN go build -o /app/dclock

ENTRYPOINT ["/app/dclock"]

