# syntax=docker/dockerfile:1.6      <-- enables $BUILDPLATFORM / $TARGETPLATFORM magic

############################  builder  ################################
ARG GO_VERSION=1.22.3              # keep in one place; change when you upgrade Go
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-bookworm AS builder

# Build-time targets supplied by Buildx
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT   # e.g. "v8" for arm64/v8

WORKDIR /src

# 1️⃣  Download deps with cache
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 2️⃣  Copy source + build for the *target* arch, but on the *host* tool-chain
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    GOOS="$TARGETOS" GOARCH="$TARGETARCH" GOARM64="$TARGETVARIANT" \
    go build -trimpath -ldflags="-s -w" -o /out/gojira ./...

############################  runtime  ################################
FROM --platform=$TARGETPLATFORM debian:bookworm-slim

# (optional) non-root user for better container hygiene
RUN useradd -r -s /sbin/nologin gojira
WORKDIR /app

COPY --from=builder /out/gojira /usr/local/bin/gojira
USER gojira

# EXPOSE 8080    # ← uncomment if your app listens here
ENTRYPOINT ["/usr/local/bin/gojira"]
