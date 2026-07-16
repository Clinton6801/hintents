# Metrics Package

This package provides Prometheus metrics for monitoring remote node health and simulation execution in the ERST daemon.

## Overview

The metrics package exports four key metrics:

1. `remote_node_last_response_timestamp_seconds` - Tracks the last successful response time from each remote node
2. `remote_node_response_total` - Counts total responses by status (success/error)
3. `remote_node_response_duration_seconds` - Histogram of response durations
4. `simulation_execution_total` - Counts total simulation executions by status

## Usage

### Recording Remote Node Responses

```go
import (
    "time"
    "github.com/dotandev/hintents/internal/metrics"
)

// Record a successful response
startTime := time.Now()
// ... make RPC call ...
duration := time.Since(startTime)
metrics.RecordRemoteNodeResponse("https://soroban-testnet.stellar.org", "testnet", true, duration)

// Record a failed response
metrics.RecordRemoteNodeResponse("https://soroban-testnet.stellar.org", "testnet", false, duration)
```

### Recording Simulation Executions

```go
// Record a successful simulation
metrics.RecordSimulationExecution(true)

// Record a failed simulation
metrics.RecordSimulationExecution(false)
```

### Exposing Metrics via HTTP

```go
import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Add metrics endpoint to your HTTP server
http.Handle("/metrics", promhttp.Handler())
```

## Testing

Run unit tests:
```bash
go test ./internal/metrics
```

Run integration tests:
```bash
go test -tags=integration ./internal/metrics
```

## Metric Details

### remote_node_last_response_timestamp_seconds

- **Type**: Gauge
- **Labels**: `node_address`, `network`
- **Updates**: Only on successful responses
- **Purpose**: Enable staleness alerting

### remote_node_response_total

- **Type**: Counter
- **Labels**: `node_address`, `network`, `status`
- **Updates**: On every response (success or error)
- **Purpose**: Track error rates and throughput

### remote_node_response_duration_seconds

- **Type**: Histogram
- **Labels**: `node_address`, `network`
- **Buckets**: `[0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10]`
- **Updates**: On every response (success or error)
- **Purpose**: Track latency and performance

### simulation_execution_total

- **Type**: Counter
- **Labels**: `status`
- **Updates**: On every simulation execution
- **Purpose**: Track overall system throughput
