# Golang Client-Server Application

## Overview

This project is a client-server application written in Go, where the server reads commands from an external RabbitMQ queue and executes them. The client sends commands to the queue. The application demonstrates the use of Go's concurrency features, including goroutines and channels.

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your machine.
- [Go](https://golang.org/dl/) installed (for local development and running the application without Docker).

```
go mod tidy
```

## Running the application

```
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

```
go run main.go -mode server -rabbitmq amqp://guest:guest@localhost:5672/ -queue commands
```

```
go run main.go -mode client -rabbitmq amqp://guest:guest@localhost:5672/ -queue commands -input commands1.txt
```

```
go run main.go -mode client -rabbitmq amqp://guest:guest@localhost:5672/ -queue commands -input commands2.txt
```

## Testing

```
go test ./...
```

## Assumptions

The order of command processing is not guaranteed. Commands are executed in parallel, so the execution order may differ from the order in which they were sent.
