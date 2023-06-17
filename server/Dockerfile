FROM golang:1.20

WORKDIR /app/server

COPY ./server .

COPY ./shared ../shared

RUN go mod download

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]