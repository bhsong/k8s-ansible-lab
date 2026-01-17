FROM golang:alpine AS builder
WORKDIR /src
COPY . .

RUN go mod init k8s-lab || true
RUN go get github.com/lib/pq
RUN go build -o myapp .

FROM alpine
WORKDIR /app
COPY --from=builder /src/myapp .
CMD ["./myapp"]