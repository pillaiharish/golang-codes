package main

import (
	"fmt"
	"sync"
)

type command struct {
	action string
	id int
	message string
	subscriber chan string
	done chan struct{}
}

// One coordinator owns the subscriber map and closes subscriber channels.
func runBus(commands <-chan command) {
	subscribers := make(map[int]chan string)
	for cmd := range commands {
		switch cmd.action {
		case "subscribe":
			subscribers[cmd.id] = cmd.subscriber
		case "publish":
			for id, ch := range subscribers {
				select {
				case ch <- cmd.message:
				default:
					fmt.Println("dropped for slow subscriber:", id)
				}
			}
		case "unsubscribe":
			if ch, ok := subscribers[cmd.id]; ok {
				close(ch)
				delete(subscribers, cmd.id)
			}
		}
		close(cmd.done)
	}
	for _, ch := range subscribers { close(ch) }
}

func main() {
	commands := make(chan command)
	go runBus(commands)
	send := func(c command) {
		c.done = make(chan struct{})
		commands <- c
		<-c.done
	}
	var wg sync.WaitGroup
	for id := 1; id <= 2; id++ {
		ch := make(chan string, 2)
		send(command{action: "subscribe", id: id, subscriber: ch})
		wg.Add(1)
		go func(subscriberID int, received <-chan string) {
			defer wg.Done()
			for msg := range received { fmt.Println(subscriberID, msg) }
		}(id, ch)
	}
	send(command{action: "publish", message: "hello"})
	send(command{action: "unsubscribe", id: 1})
	send(command{action: "publish", message: "again"})
	close(commands)
	wg.Wait()
}
