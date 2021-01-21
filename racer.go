package racer

import "fmt"

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

// Infructrucure

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

// func ping(timeout time.Duration) func(string) chan bool {
// 	return func(domain string) chan bool {
// 		return realPing(domain, timeout)
// 	}
// }
