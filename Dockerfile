FROM golang:1.23.2-alpine AS builder

WORKDIR /remindings

RUN apk add --no-cache tzdata

# dependencies
COPY ["Makefile", "go.mod", "./"]
RUN go mod download

COPY . .

# build
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app ./cmd/remindings

FROM alpine:3.20

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder --chown=appuser:appgroup --chmod=755 /app /app

USER appuser

EXPOSE 8080

CMD ["app"]
