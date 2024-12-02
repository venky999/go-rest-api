# ---- Build Phase ----
FROM golang:1.23 AS builder
COPY . /app
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /app/go-rest-api cmd/main.go

# ---- Deploy Phase ----
FROM gcr.io/distroless/static:nonroot
USER nonroot:nonroot
COPY --from=builder  --chown=nonroot:nonroot /app/go-rest-api /app/go-rest-api
ENTRYPOINT ["/app/go-rest-api"]