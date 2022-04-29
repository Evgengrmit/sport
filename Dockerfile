FROM golang:1.18.1

ADD . /code
WORKDIR /code
RUN cd /code
RUN go mod download

CMD ["go", "run", "cmd/server/main.go"]