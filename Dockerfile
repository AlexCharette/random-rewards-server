FROM golang:alpine AS build

WORKDIR /build

COPY go.mod .
COPY . .

RUN go build -o random-rewards main.go

FROM alpine

LABEL org.opencontainers.image.source=https://github.com/AlexCharette/random-rewards-server

WORKDIR /build

COPY --from=build /build/random-rewards /build/random-rewards

EXPOSE 8080

CMD [ "/build/random-rewards" ]