package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Total time: %.4fs\n", time.Since(start).Seconds())
	}()

	items := []string{
		"http://google.com",
		"http://python.org",
		"http://ruby-lang.org",
		"http://golang.org",
	}

	c := make(chan string)

	for _, item := range items {
		go count(item, c)
	}

	timeout := time.After(500 * time.Millisecond)
	for _ = range items {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-timeout:
			return
		}
	}
}

func count(url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(ioutil.Discard, r.Body)
	if err != nil {
		panic(err)
	}
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s %d [%.4fs]\n", url, n, dt)
}
