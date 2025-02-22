package main

import (
	"encoding/json"
	"fmt"
	"time"

	beanstalk "github.com/beanstalkd/go-beanstalk"
)

type message struct {
	UserID    uint `json:"userId"`
	ContentID uint `json:"contentId"`
	PageRead  int  `json:"pageRead"`
}

type queueHandler struct {
	conn *beanstalk.Conn
}

func setQueueHandler(host string) (*queueHandler, error) {
	conn, err := beanstalk.Dial("tcp", host)
	return &queueHandler{
		conn,
	}, err
}

func (qh *queueHandler) readFromTube(name string, output chan message) {
	tube := beanstalk.NewTube(qh.conn, name)
	for {
		time.Sleep(1 * time.Second) // hacky, shouldn't do this in real world
		id, body, err := tube.PeekReady()
		if id == 0 {
			continue
		}
		fmt.Printf("Queue [%s]. Read id %d, with: %s\n", name, id, string(body))
		qh.conn.Delete(id) // delete before unmarshalling; if that fails, queue will be blocked
		msg := message{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			continue
		}
		output <- msg
	}
}
