package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_task/common"
	"go_task/queue"
	"log"
	"os"
)

type Client struct {
	queue queue.Queue
}

func NewClient(q queue.Queue) *Client {
	return &Client{queue: q}
}

func (c *Client) Run(inputFile string) {
	var commands []common.Command

	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			var cmd common.Command
			if err := json.Unmarshal([]byte(line), &cmd); err != nil {
				log.Printf("Error unmarshaling command: %v", err)
				continue
			}
			commands = append(commands, cmd)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter commands:")
		for scanner.Scan() {
			line := scanner.Text()
			var cmd common.Command
			if err := json.Unmarshal([]byte(line), &cmd); err != nil {
				log.Printf("Error unmarshaling command: %v", err)
				continue
			}
			commands = append(commands, cmd)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading stdin: %v", err)
		}
	}

	for _, cmd := range commands {
		msg, err := json.Marshal(cmd)
		if err != nil {
			log.Printf("Error marshaling command: %v", err)
			continue
		}

		if err := c.queue.SendMessage(string(msg)); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}