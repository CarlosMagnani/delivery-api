FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download


COPY *.go ./
RUN go build ./main.go

EXPOSE 4000

CMD [ "/delivery-api/cmd/delivery" ]