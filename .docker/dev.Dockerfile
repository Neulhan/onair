FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN go get github.com/cosmtrek/air

RUN go get github.com/pilu/fresh

ENTRYPOINT ["air"]
