package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

func readFromFile(nameFile string) {

	file, err := os.Open(nameFile)
	if err != nil {
		log.Fatal(err)
	}

	text := bytes.Buffer{}
	fileText := bufio.NewScanner(file)
	for fileText.Scan() {
		text.WriteString(fileText.Text())
	}
	fmt.Println(text.String())
	wg.Done()
}

func main() {

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		nameFile := strconv.Itoa(i) + ".txt"
		go readFromFile(nameFile)
	}
	wg.Wait()

}
