# GO REST API

## Introduction
REST API written in go with postgres database

## Features
-   Insert transaction into a DB using UUID
-   Dockerized and Working locally using docker compose
-   Zap based uniform logging
-   Security Headers
-   Health Check Endpoint
-   Metrics Endpoint
-   Input Validation
-   Uses Distroless nonroot image
-   [TODO] Secure CI/CD Pipeline
-   [TODO] Cloud Infrasture Automation and Security

## API
Available APIs are
```
# Create Transaction

POST /api/transaction --data
'{"transactionId":"$UUID","amount": "$AMOUNT","timestamp":"$TIME"}'
```
```
# Get metrics

GET /metrics
```
```
# Get health info

GET /health
```

## Installation

### 1. Build
make build

### 2. Publish
make push

### 3. Run
make run

### 4. Stop
make stop

### 5. Cleanup
make cleanup

### 6. Testing

```bash
# Transaction Request
 curl -X POST http://localhost:8080/api/transaction \
--data '{"transactionId":"9577A425-C385-4E35-9D35-409898DE4072","amount": "199.99","timestamp":"2024-12-02T08:045:15Z"}'

# Response
{"Success"}             
```
## Networking

| Port | Description |
|-----------|-------------|
| `8080`    | API Listening Port |
| `5432`    | Posgres Listening Port |
