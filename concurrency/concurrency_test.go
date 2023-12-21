package concurrency

import (
	"reflect"
	"testing"
	"time"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wb WebsiteChecker, urls []string) map[string]bool {
	rc := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wb(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		rc[r.string] = r.bool
	}
	return rc
}

func mockWebsiteChecker(url string) bool {
	if url == "https://bad.com" {
		return false
	} else {
		return true
	}
}

func slowChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsiteChecker(t *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "https:/a.com"
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		CheckWebsite(slowChecker, urls)
	}
}

func TestWebSiteChecker(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://apple.com",
		"https://bad.com",
	}
	want := map[string]bool{
		"https://google.com": true,
		"https://apple.com":  true,
		"https://bad.com":    true,
	}
	got := CheckWebsite(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v but got %v", want, got)
	}

}
