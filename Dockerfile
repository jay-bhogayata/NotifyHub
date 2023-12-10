FROM golang:1.21.5-alpine3.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o notification-api .

FROM alpine:3.14.2
WORKDIR /
COPY --from=builder /app/notification-api notification-api
EXPOSE 8080
CMD [ "./notification-api" ]