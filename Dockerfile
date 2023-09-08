FROM golang:1.21-alpine

COPY . /booking-app/

WORKDIR /booking-app/

RUN go build -o /booking-app/bin/bookingapp

CMD [ "/booking-app/bin/bookingapp" ]
