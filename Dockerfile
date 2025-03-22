FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o microservice .

EXPOSE 9003

CMD ["./microservice"]

