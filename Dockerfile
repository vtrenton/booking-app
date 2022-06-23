FROM golang:1.18-alpine

WORKDIR /booking-app

COPY go.mod ./
COPY helper/ ./
COPY main.go ./

RUN go build -o /booking-app

CMD [ "booking-app" ]