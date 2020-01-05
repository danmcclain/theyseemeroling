FROM golang:1.13.5 AS builder
WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o theyseemeroling cmd/theyseemeroling/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder app/theyseemeroling .
CMD ["./theyseemeroling"] 
