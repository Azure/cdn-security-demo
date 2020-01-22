FROM golang:1.13.6-alpine as builder

# Setup build environemnt.
ENV CGO_ENABLED=0
WORKDIR /src

# Download dependencies first so they stay cached when changing local source.
COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
RUN go mod download

# Copy local source code and build.
COPY main.go /src
RUN go build -o /usr/bin/origin .

# Copy built binary into a "scratch" container
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/bin/origin /usr/bin/origin
ENTRYPOINT ["/usr/bin/origin"]