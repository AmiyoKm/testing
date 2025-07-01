package concurrency

import (
	"fmt"
	"reflect"
	"testing"
)

func MockWebsiteChecker(url string) bool {
	return url != "http://gtfo.com"
}
func TestCheckWebsites(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://yahoo.com",
		"http://gtfo.com",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://yahoo.com":  true,
		"http://gtfo.com":   false,
	}

	got := CheckWebsites(MockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v , got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = fmt.Sprintf("http://google.com/%d", i)
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
