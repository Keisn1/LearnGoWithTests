package main

import (
	"net/http"
	"time"
)

const (
	ErrTimeout = RequestErr("Took longer than 10 seconds")
)

type RequestErr string

func (rError RequestErr) Error() string {
	return string(rError)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(ch chan struct{}) {
		http.Get(url)
		close(ch)
	}(ch)

	return ch

}

var tenSecondTimout = 10 * time.Second

func Racer(u1, u2 string) (string, error) {
	return ConfigurableRacer(u1, u2, tenSecondTimout)
}

func ConfigurableRacer(u1, u2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(u1):
		return u1, nil
	case <-ping(u2):
		return u2, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}
