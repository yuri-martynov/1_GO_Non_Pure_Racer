package racer

import (
	"testing"
	"time"
)

func Test_A_is_fastest(t *testing.T) {

	ping := func(url string, timeout time.Duration) chan bool {
		ch := make(chan bool)
		go func() {
			if url == "A" {
				ch <- true
			}
		}()
		return ch
	}

	domain, _ := racer("A", "B", ping, time.Second)
	if domain != "A" {
		t.Error("A should be the fastes")
	}
}

func Test_B_is_fastest(t *testing.T) {
	ping := func(url string, timeout time.Duration) chan bool {
		ch := make(chan bool)
		go func() {
			if url == "B" {
				ch <- true
			}
		}()
		return ch
	}

	domain, _ := racer("A", "B", ping, time.Second)
	if domain != "B" {
		t.Error("B should be the fastes")
	}
}

func Test_Yandex_or_Google(t *testing.T) {
	x, _ := racer("yandex.ru", "google.com", ping, 10*time.Second)
	if x != "yandex.ru" && x != "google.com" {
		t.Error("Smth should be the fastes")
	}
}

func Test_Error(t *testing.T) {
	ping := func(url string, timeout time.Duration) chan bool {
		ch := make(chan bool)
		go func() {
			ch <- false
		}()
		return ch
	}
	_, err := racer("A", "B", ping, time.Second)
	if err == nil {
		t.Error("Shuld be failed")
	}
}
