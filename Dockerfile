FROM golang:1.18-alpine

WORKDIR /app

COPY ./server/go.mod ./
COPY ./server/go.sum ./
RUN go mod download

COPY ./server/*.go ./

RUN go build -o /email


CMD [ "/email" ]