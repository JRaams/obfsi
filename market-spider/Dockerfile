# Build stage
FROM golang:1.15-alpine as builder

RUN apk update
RUN apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    git \
    musl-dev \
    ca-certificates

RUN update-ca-certificates

RUN mkdir build
COPY go.mod go.sum ./build/
RUN cd ./build/ && go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . /build
WORKDIR /build

RUN go build -ldflags '-extldflags "-static"' -o market-spider .
RUN adduser -S -D -H -h /build market-spider
USER market-spider

# Run stage
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=builder /build/market-spider /app/
CMD [ "./market-spider" ]