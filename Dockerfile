FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o builded

CMD ["./builded"]

EXPOSE 8080

ENTRYPOINT ["./builded"]