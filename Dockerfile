FROM golang:1.17

RUN mkdir /app

ADD . /app

WORKDIR /app

ENV PORT=8080 \
    DB_HOST=ec2-3-224-157-224.compute-1.amazonaws.com \
    DB_PORT=5432  \
    DB_NAME=dbiqfdis1j5q6g \
    DB_USER=xzmbprjsejsazz \
    DB_PASSWORD=1b2e39d2b1b6d7098cf2a756a4706276817729749cde329086b2772ee1f9d74a

RUN go mod init car-api

RUN go mod tidy

RUN go get -u github.com/swaggo/swag/cmd/swag@v1.7.0

COPY . .

RUN swag init

RUN go mod tidy

RUN go build main.go

CMD ["/app/main"]