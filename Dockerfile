FROM golang:1.17 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o consumer .

FROM alpine:3.12

WORKDIR /app

COPY --from=build /app/consumer .

CMD ["./consumer"]


