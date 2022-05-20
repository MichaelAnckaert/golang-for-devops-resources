# Build image
FROM golang:1.18-bullseye AS builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-server .

# Deploy image
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/hello-server .
CMD /app/hello-server .