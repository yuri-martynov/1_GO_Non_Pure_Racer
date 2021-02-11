package racer

import (
	"fmt"
	"testing"
	"time"
)

func Test_A_is_fastest(t *testing.T) {

	ping := func(url string) chan error {
		ch := make(chan error)
		go func() {
			if url == "A" {
				ch <- nil
			}
		}()
		return ch
	}

	domain, _ := racer("A", "B", ping)
	if domain != "A" {
		t.Error("A should be the fastes")
	}
}

func Test_B_is_fastest(t *testing.T) {

	ping := func(url string) chan error {
		ch := make(chan error)
		go func() {
			if url == "B" {
				ch <- nil
			}
		}()
		return ch
	}

	domain, _ := racer("A", "B", ping)

	if domain != "B" {
		t.Error("B should be the fastes")
	}
}

func Test_B_if_A_has_error(t *testing.T) {
	// setup / given
	ping := func(url string) chan error {
		ch := make(chan error)
		go func() {
			if url == "A" {
				ch <- fmt.Errorf("A has an error")
			} else {
				time.Sleep(100 * time.Millisecond)
				ch <- nil
			}
		}()
		return ch
	}

	// test / when
	domain, _ := racer("A", "B", ping)

	// check / then
	if domain == "B" {
		return
	}

	t.Error("A has an error")
}

func Test_A_if_B_has_error(t *testing.T) {
	// setup / given
	ping := func(url string) chan error {
		ch := make(chan error)
		go func() {
			if url == "B" {
				ch <- fmt.Errorf("B has an error")
			} else {
				time.Sleep(100 * time.Millisecond)
				ch <- nil
			}
		}()
		return ch
	}

	// test / when
	domain, _ := racer("A", "B", ping)

	// check / then
	if domain == "A" {
		return
	}

	t.Error("B has an error")
}

func Test_all_pings_failed(t *testing.T) {
	// setup / given
	ping := func(url string) chan error {
		ch := make(chan error)
		go func() {
			ch <- fmt.Errorf("B has an error")
		}()
		return ch
	}

	// test / when
	_, err := racer("A", "B", ping)

	// check / then
	if err != nil {
		return
	}

	t.Error("All pings have an error")
}

func Test_get_is_working(t *testing.T) {
	// setup / given

	// test / when
	_, err := racer("https://yandex.ru", "https://google.com", get)

	// check / then
	if err == nil {
		return
	}

	t.Error("get return error")
}

func Test_adapter_returns_timeout(t *testing.T) {
	// setup
	ping := func(url string) chan error {
		ch := make(chan error)
		go func() {
			time.Sleep(time.Second)
			ch <- nil
		}()
		return ch
	}

	pingWithTimeout := withTimeout(ping, time.Millisecond)

	// test
	err := <-pingWithTimeout("A")

	// check
	if err == errTimeout {
		return
	}

	t.Error("should be timed out")

}

func Test_adapter_returns_ping_nil(t *testing.T) {
	// setup
	ping := func(url string) chan error {
		ch := make(chan error)
		go func() {
			time.Sleep(time.Millisecond)
			ch <- nil
		}()
		return ch
	}

	pingWithTimeout := withTimeout(ping, time.Second)

	// test
	err := <-pingWithTimeout("A")

	// check
	if err == nil {
		return
	}

	t.Error("error should be nil")

}

func Test_Get_returns_Timeout(t *testing.T) {
	// setup

	pingWithTimeout := getWithTimeout(time.Microsecond)

	// test
	err := <-pingWithTimeout("http://ya.ru")

	// check
	if err == errTimeout {
		return
	}

	t.Error("error should be timeout")

}

func Test_Get_returns_Reply(t *testing.T) {
	// setup

	pingWithTimeout := getWithTimeout(time.Second * 30)

	// test
	err := <-pingWithTimeout("http://ya.ru")

	// check
	if err == nil {
		return
	}

	t.Error("error should be nil")

}
