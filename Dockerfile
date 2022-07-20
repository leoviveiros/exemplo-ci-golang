FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]