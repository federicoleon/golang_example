package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	SORT_URL = "http://localhost:8080/sort"
	LIMIT    = 100000
)

func main() {
	coreNum := runtime.NumCPU()
	runtime.GOMAXPROCS(coreNum)

	c := make(chan string)
	defer close(c)

	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < LIMIT; i++ {
		wg.Add(1)
		go getOrder(c)
	}

	for i := 0; i < LIMIT; i++ {
		wg.Done()
		fmt.Println(<-c)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("%d took %s", LIMIT, elapsed))
}

func getOrder(c chan string) {
	var json = []byte(`{"original":[4,6,2,7,3,8,1,9,0]}`)
	req, err := http.NewRequest(http.MethodPost, SORT_URL, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(20 * time.Millisecond),
	}
	resp, err := client.Do(req)
	if err != nil {
		c <- err.Error()
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- err.Error()
		return
	}

	result := string(body)
	c <- result
}
