package racer

import "fmt"

// LOGIC

type funcPing = func(endpoint string) chan error

func racer(a string, b string, ping funcPing) (error, string) {

	aChan := ping(a)
	bChan := ping(b)

	select {
	case err := <-aChan:
		if err == nil {
			return nil, a
		}

		err = <-bChan
		if err == nil {
			return nil, b
		}

	case err := <-bChan:
		if err == nil {
			return nil, b
		}

		err = <-aChan
		if err == nil {
			return nil, a
		}
	}

	return fmt.Errorf("all pings failed"), ""
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
