FROM core.harbor.k8s.devim.team/proxy/library/golang:1.17-alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /app
COPY . .
RUN go build \
    -ldflags "-s -w -extldflags '-static'" \
    -mod=vendor \
    -o /example \
    ./example

FROM scratch
COPY --from=builder /example /example
ENTRYPOINT ["/example"]
