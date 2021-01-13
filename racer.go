package racer

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func racer(a string, b string, ping func(string, time.Duration) chan bool, timeout time.Duration) (string, error) {

	select {
	case chA := <-ping(a, timeout):
		if chA {
			return a, nil
		} else {
			return "", fmt.Errorf("ping timeout")
		}
	case chB := <-ping(b, timeout):
		if chB {
			return b, nil
		} else {
			return "", fmt.Errorf("ping timeout")
		}
	}
}

func ping(domain string, timeout time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		var err = exec.Command("ping", domain, "-n 1", "-w "+strconv.FormatInt(timeout.Milliseconds(), 10)).Run()
		if err != nil {
			ch <- false
		} else {
			ch <- true
		}
	}()
	return ch
}
