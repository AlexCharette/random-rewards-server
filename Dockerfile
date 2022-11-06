FROM golang:1.19.3

LABEL org.opencontainers.image.source=https://github.com/AlexCharette/random-rewards-server

# TODO: Set a user

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /random-rewards

EXPOSE 8080

CMD [ "/random-rewards" ]

