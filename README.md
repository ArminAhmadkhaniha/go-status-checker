# Go Status Checker

A high-performance network monitor demonstrating the evolution from sequential processing to concurrent microservices in Go.

## Project Structure

This repository contains three versions of the application to demonstrate different engineering patterns:

### 1. Sequential (`/01-sequential`)
- Checks websites one by one.
- **Performance:** Slow (~1.5s).
- **Use Case:** Simple scripts where speed doesn't matter.
- **Run:** `go run 01-sequential/main.go`

### 2. Concurrent Finite (`/02-concurrent`)
- Uses **Goroutines** and `sync.WaitGroup` to check all sites simultaneously.
- **Performance:** Fast (~0.3s).
- **Use Case:** Batch processing, cron jobs, lambda functions.
- **Run:** `go run 02-concurrent/main.go`

### 3. Infinite Service (`/03-infinite-service`)
- The final architecture. Runs as a daemon using **Channels** for a feedback loop.
- **Behavior:** Checks status, waits 5 seconds, repeats forever.
- **Use Case:** Real-time monitoring, microservices, 5G network probes.
- **Run:** `go run 03-infinite-service/main.go`

## Technology
- **Language:** Go (Golang)
- **Key Concepts:** Goroutines, Channels, WaitGroups, Non-blocking I/O