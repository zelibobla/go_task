package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"go_task/common"
	"go_task/orderedmap"
	"go_task/queue"
)

type Server struct {
	queue      queue.Queue
	data       *orderedmap.OrderedMap
	outputFile string
	wg         sync.WaitGroup
}

func NewServer(q queue.Queue, outputFile string) *Server {
	return &Server{
		queue:      q,
		data:       orderedmap.NewOrderedMap(),
		outputFile: outputFile,
	}
}

func (s *Server) Start() {
	msgs, err := s.queue.ReceiveMessages()
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}

	log.Println("Connected to RabbitMQ. Listening...")

	for msg := range msgs {
		var cmd common.Command
		if err := json.Unmarshal(msg.Body, &cmd); err != nil {
			log.Printf("Error unmarshaling command: %v", err)
			continue
		}

		s.wg.Add(1)
		go s.executeCommand(cmd)
	}
}

func (s *Server) executeCommand(cmd common.Command) {
	defer s.wg.Done()

	switch cmd.Action {
	case "addItem":
		log.Printf("AddItem: %s", cmd.Key)
		s.data.Add(cmd.Key, cmd.Value)
	case "deleteItem":
		log.Printf("DeleteItem: %s", cmd.Key)
		s.data.Delete(cmd.Key)
	case "getItem":
		if value, exists := s.data.Get(cmd.Key); exists {
			log.Printf("GetItem: %s = %s", cmd.Key, value)
			s.writeToFile(fmt.Sprintf("GetItem: %s = %s\n", cmd.Key, value))
		} else {
			log.Printf("GetItem: %s not found", cmd.Key)
			s.writeToFile(fmt.Sprintf("GetItem: %s not found\n", cmd.Key))
		}
	case "getAllItems":
		allItems := s.data.GetAll()
		for k, v := range allItems {
			log.Printf("GetAllItems: %s = %s", k, v)
			s.writeToFile(fmt.Sprintf("GetAllItems: %s = %s\n", k, v))
		}
	}
}

func (s *Server) writeToFile(data string) {
	file, err := os.OpenFile(s.outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}

func (s *Server) Stop() {
	s.wg.Wait()
}
