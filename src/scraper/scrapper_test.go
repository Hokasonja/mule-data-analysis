package scrapper

import (
	"testing"
)

func TestScrape(t *testing.T) {
	// 스크래핑 함수에 대한 테스트 코드
	testURL := "https://example.com"
	data, err := Scrape(testURL)
	if err != nil {
		t.Errorf("Scrape() error: %v", err)
	}
	if len(data) == 0 {
		t.Error("Scraped data is empty")
	}
}
