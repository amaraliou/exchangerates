FROM golang:stretch as builder
COPY . /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

RUN go build -o /curve-exchangerates .

FROM heroku/heroku:16
COPY --from=builder /curve-exchangerates /curve-exchangerates
CMD ["/curve-exchangerates"]