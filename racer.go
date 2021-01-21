package racer

// LOGIC

func racer(a string, b string, ping func(string) chan bool) string {

	select {
	case <-ping(a):
		return a

	case <-ping(b):
		return b
	}
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
