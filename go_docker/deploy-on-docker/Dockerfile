FROM golang:1.18.6

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main ./main.go

EXPOSE 8080

CMD /app/main
