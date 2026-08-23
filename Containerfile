# Build stage
FROM docker.io/library/golang:1.26-alpine AS builder

WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /bin/api ./cmd/api

# Runtime stage
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /bin/api /api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/api"]
