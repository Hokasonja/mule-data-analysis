package scrapper

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Scrape 함수는 주어진 URL에서 데이터를 스크래핑합니다.
func Scrape(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
