package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Website string

func (w Website) wait(d time.Duration) {
	time.Sleep(d * time.Second)
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers. Returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Errorf("Got an unexpected error")
		}
		if got != want {
			t.Errorf("WebsiteRacer() = \"%v\"; want \"%v\"", got, want)
		}
	})
	t.Run("compares speed of servers. Returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(2 * time.Millisecond)
		fastServer := makeDelayedServer(2 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, time.Millisecond)

		if err != ErrTimeout {
			t.Errorf("Did not get the right error message")
		}
	})
}
