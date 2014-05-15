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
	for i := 0; i < len(items); i++ {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-timeout:
			fmt.Printf("Total time: %.5fs\n", time.Since(start).Seconds())
			return
		}
	}

	fmt.Printf("Total time: %.2fs\n", time.Since(start).Seconds())
}

func count(url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s\n", url, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", url, n, dt)
}
