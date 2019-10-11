FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add --no-cache git
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/janosik834/Calculate/calcDocker/calculateP
RUN go build -o main
CMD ["/app/main"]
