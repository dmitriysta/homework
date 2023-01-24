package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	NewWorker(ctx)
}

func NewWorker(ctx context.Context) Worker {
	w := Worker{
		buffer: make([]string, 0),
	}
	go w.scanData()
	go w.timerAddData(ctx)
	return w
}

type Worker struct {
	buffer []string
	mutex  sync.Mutex
}

// считывает инфо с консоли и добавляет в буффер
func (w *Worker) scanData() {
	var data string
	fmt.Println("Add your information:")
	fmt.Scan(&data)
	w.mutex.Lock()
	w.buffer = append(w.buffer, data)
	w.mutex.Unlock()
}

// добавляет инфо в файл
func (w *Worker) addData() {
	var temp []string

	w.mutex.Lock()
	temp = w.buffer
	w.buffer = make([]string, 0)
	w.mutex.Unlock()

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Unable to read file:", err)
	}
	for _, value := range temp {
		file.WriteString(value)
	}
}

// добавление таймера (1 минута)
func (w *Worker) timerAddData(ctx context.Context) {
	timer := time.NewTimer(time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			w.addData()
		}
	}
}
