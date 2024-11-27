# Build STAGE
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# RUN STAGE
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY app.env .
EXPOSE 8080
CMD ["/app/main"]