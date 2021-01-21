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

	_, domain := racer("A", "B", ping)
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

	_, domain := racer("A", "B", ping)
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
	_, domain := racer("A", "B", ping)

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
	_, domain := racer("A", "B", ping)

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
	err, _ := racer("A", "B", ping)

	// check / then
	if err != nil {
		return
	}

	t.Error("All pings have an error")
}
