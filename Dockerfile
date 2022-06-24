FROM golang:1.18-alpine

WORKDIR /usr/local/go/src/booking-app/

COPY helper/ ./helper
COPY main.go ./
COPY go.mod ./

RUN go build -o /booking-app

CMD [ "/booking-app" ]