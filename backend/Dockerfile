FROM golang:1.22-alpine

RUN apk update && apk upgrade

WORKDIR /app

COPY go.mod go.sum ./

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .

EXPOSE 7777

CMD ["./main"]
