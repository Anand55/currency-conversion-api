FROM golang:1.17

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV FIXER_KEY d244f7d57ffda45bc7b3b39e1ae75d0d
ENV REDIS_ADDR redis:6379

RUN go build -o conversion-api ./cmd 

CMD ["./conversion-api"]