FROM golang:1.17

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV FIXER_KEY 1fe73a0bf57fa882382b2b7d18067c05

RUN go build -o conversion-api ./cmd 

CMD ["./conversion-api"]