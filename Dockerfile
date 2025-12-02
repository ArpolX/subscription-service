FROM golang:1.24.10-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o subscription_service ./cmd/app

FROM alpine:latest
COPY --from=builder /app/subscription_service /subscription_service
CMD [ "/subscription_service" ]