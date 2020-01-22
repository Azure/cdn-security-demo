FROM golang:1.13.6-alpine as builder

ENV CGO_ENABLED=0
WORKDIR /src
COPY . /src/origin
RUN go build -o /usr/bin/origin ./origin

FROM scratch

COPY --from=builder /usr/bin/origin /usr/bin/origin
ENTRYPOINT ["/usr/bin/origin"]