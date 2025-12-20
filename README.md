# NATS by Examples

## Description

This repository contains a collection of practical examples demonstrating various messaging patterns using NATS, implemented in Go. NATS is a lightweight, high-performance messaging system for cloud-native applications, microservices, and distributed systems. The examples cover fundamental to advanced concepts, making it an educational resource for learning NATS.

This project was developed as part of my portfolio to showcase skills in Go programming, asynchronous messaging, distributed systems, and containerization with Docker.

## Examples

- [x] 1. Basic Connection: Establishes a simple connection to a NATS server.
- [x] 2. Publish-Subscribe Pattern: Demonstrates broadcasting messages to multiple subscribers.
- [x] 3. Pub-Sub with JSON: Handles structured data (JSON) in publish-subscribe scenarios.
- [x] 4. Wildcard Subscriptions: Shows how to use wildcards for flexible topic matching.
- [x] 5. Request-Reply Pattern: Implements synchronous request-response communication.
- [x] 6. Multiple Repliers (Load Balancing)
- [ ] 7. Queue Groups
- [ ] 8. JetStream with Go
- [ ] 9. Retry with Exponential Backoff
- [ ] 10. Circuit Breaker Pattern

## Technologies Used

- **Go**: Primary programming language (version 1.21 or higher).
- **NATS**: Messaging system (via the `nats.go` library).
- **Docker**: For running the NATS server in a containerized environment.
- Other dependencies: See `go.mod` for the full list, including `nats.go` and related utilities.

## Installation

### Prerequisites

- Go 1.21 or higher installed.
- Docker and Docker Compose (for running the NATS server).

### From Source

1. Clone the repository:
   ```
   git clone https://github.com/arashalaei/nats-by-examples.git
   cd nats-by-examples
   ```
2. Install dependencies:
   ```
   go mod download
   ```

## Usage

### Running the NATS Server

Use Docker Compose to start a local NATS server:

```
docker-compose up -d
```

This will run NATS on `localhost:4222`.

### Running Examples

Each example is in its own directory and can be run independently. Navigate to the directory and execute the Go file.

For example, to run the basic connection:

```
cd 1.basic-connection
go run main.go
```

Similarly for other examples. Some examples may require running multiple terminals (e.g., one for publisher, one for subscriber).

## Project Structure

- `1.basic-connection/`: Basic NATS connection example.
- `2.pub-sub-pattern/`: Publish-subscribe pattern.
- `3.pub-sub-with-json/`: Pub-sub with JSON payloads.
- `4.wildcard-subscriptions/`: Wildcard topic subscriptions.
- `5.request-replay-pattern/`: Request-reply pattern (note: likely a typo for request-reply).
- `docker-compose.yml`: Docker configuration for NATS server.
- `go.mod` and `go.sum`: Go module files.

## Contributing

Contributions are welcome! Feel free to add new examples, fix bugs, or improve documentation. Fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details (note: add a LICENSE file if not present).

## Contact

For questions or collaboration, reach out via GitHub issues or my portfolio site.
