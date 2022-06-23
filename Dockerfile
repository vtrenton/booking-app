FROM golang:1.18-alpine

WORKDIR /usr/local/go/src/booking-app/

COPY helper/ ./
COPY main.go ./
COPY go.mod ./

RUN go build -o /booking-app-bin

CMD [ "/booking-app-bin" ]