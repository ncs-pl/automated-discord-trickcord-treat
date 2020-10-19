FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM alpine:latest
COPY --from=builder /build/main /app/
COPY --from=builder /build/.env /app/
WORKDIR /app
CMD ["./main"]
