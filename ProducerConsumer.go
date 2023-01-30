package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Message struct {
	mu     sync.Mutex
	buffer []string
}

func (m *Message) consumer() {
	m.mu.Lock()
	defer m.mu.Unlock()
	fmt.Println(m.buffer[0])
	m.buffer = m.buffer[1:]
}

func (m *Message) producer(msg string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.buffer = append(m.buffer, msg)
}

func main() {
	m := Message{}
	for i := 1; i <= 100; i++ {
		go m.producer(strconv.Itoa(i))
	}
	for len(m.buffer) > 0 {
		m.consumer()
	}
}
