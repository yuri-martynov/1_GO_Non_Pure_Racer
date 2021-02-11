package racer

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// LOGIC

type funcPing = func(endpoint string) chan error

func racer(a string, b string, ping funcPing) (string, error) {

	aChan := ping(a)
	bChan := ping(b)

	select {
	case err := <-aChan:
		if err == nil {
			return a, nil
		}

		err = <-bChan
		if err == nil {
			return b, nil
		}

	case err := <-bChan:
		if err == nil {
			return b, nil
		}

		err = <-aChan
		if err == nil {
			return a, nil
		}
	}

	return "", fmt.Errorf("all pings failed")
}

var errTimeout = fmt.Errorf("timeout")

func withTimeout(ping funcPing, timeout time.Duration) funcPing {
	return func(url string) chan error {
		ch := make(chan error)
		go func() {
			select {
			case <-ping(url):
				ch <- nil
			case <-time.After(timeout):
				ch <- errTimeout
			}
		}()

		return ch
	}
}

// Infructrucure

func get(url string) chan error {
	ch := make(chan error)
	go func() {
		_, err := http.Get(url)
		ch <- err
	}()
	return ch
}

func getWithTimeout(timeOut time.Duration) funcPing {
	return func(endpoint string) chan error {
		ch := make(chan error)

		go func() {
			client := http.Client{
				Timeout: timeOut,
			}
			_, err := client.Get(endpoint)
			if os.IsTimeout(err) {
				ch <- errTimeout
			} else {
				ch <- err
			}
		}()

		return ch
	}
}

// func realPing(domain string, timeout time.Duration) chan bool {
// 	ch := make(chan bool)
// 	go func() {
// 		var err = exec.Command("ping", domain, "-n 1", "-w "+strconv.FormatInt(timeout.Milliseconds(), 10)).Run()
// 		if err != nil {
// 			ch <- false

// 		} else {
// 			ch <- true
// 		}
// 	}()
// 	return ch
// }
