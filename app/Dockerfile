FROM golang:alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o myapp .

FROM alpine
WORKDIR /app
COPY --from=builder /src/myapp .
CMD ["./myapp"]