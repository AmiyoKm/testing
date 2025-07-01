package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("returns the url of the fastest server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want { 
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if server takes more than 10 seconds to response", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastServer := makeDelayedServer(12 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		_, err := ConfigurableRacer(slowUrl, fastUrl, 0*time.Second)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
