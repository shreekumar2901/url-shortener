FROM golang:alpine as build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o shortener .

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=build /app/shortener /app/

WORKDIR /app

EXPOSE 3000

CMD ["./shortener"]

