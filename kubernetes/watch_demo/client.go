package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	url = "http://localhost:8000/watch"
)

func main() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("transfer encode: %v\n", resp.TransferEncoding)

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			fmt.Print(line)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
