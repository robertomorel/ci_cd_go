FROM golang:1.18

WORKDIR /app

COPY . .

RUN go build -o math

CMD ["./math"]