package main

import (
	"fmt"
	"log"

	"github.com/Hokasonja/mule-data-analysis/src/scrapper"
	"github.com/spf13/viper"
)

type Config struct {
	ScarpingTargetUrl string
}

func main() {
	// 설정 파일 이름 및 경로 설정
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	var config Config
	// 설정 파일의 내용을 구조체로 언마샬링합니다.
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
	}

	data, err := scrapper.Scrape(config.ScarpingTargetUrl)
	if err != nil {
		log.Fatalf("Scraping error: %v", err)
	}

	fmt.Println("Scrapped data:")
	fmt.Println(data)
}
