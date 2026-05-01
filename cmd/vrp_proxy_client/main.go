package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func send(id string, wg *sync.WaitGroup) {
	defer wg.Done()

	req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/transfer", nil)
	req.Header.Set("X-Mutation-ID", id)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("=== sending concurrent requests ===")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go send("payment-001", &wg)
	}

	wg.Wait()
}