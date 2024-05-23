package main

import (
	"flag"
	"go_task/client"
	"go_task/queue"
	"go_task/server"
	"log"
)

func main() {
	mode := flag.String("mode", "server", "Mode to run: server or client")
	rabbitMQURL := flag.String("rabbitmq", "amqp://guest:guest@localhost:5672/", "RabbitMQ URL")
	queueName := flag.String("queue", "commands", "Queue name")
	outputFile := flag.String("output", "output.txt", "Output file for server")
	inputFile := flag.String("input", "", "Input file with commands for client") 
	flag.Parse()

	rmq, err := queue.NewRabbitMQ(*rabbitMQURL, *queueName)
	if err != nil {
		log.Fatalf("Error creating RabbitMQ: %v", err)
	}
	defer rmq.Close()

	if *mode == "server" {
		srv := server.NewServer(rmq, *outputFile)
		srv.Start()
	} else if *mode == "client" {
		cli := client.NewClient(rmq)
		cli.Run(*inputFile)
	} else {
		log.Fatalf("Unknown mode: %s", *mode)
	}
}
