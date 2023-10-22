FROM golang:1.21 AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 9001
ENTRYPOINT ["./main"]