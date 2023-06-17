FROM golang:1.20

WORKDIR /app/client

COPY ./client .

COPY ./shared ../shared

RUN go mod download 

RUN go build -o client .

EXPOSE 8080

CMD ["./client"]