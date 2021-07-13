package scraper

import (
	"fmt"

	"github.com/badoux/goscraper"
)

func GetImg(url *string) <-chan interface{} {
	res := make(chan interface{})
	go func() {
		defer func() {
			recover()
		}()
		s, err := goscraper.Scrape(*url, 5)
		if err != nil {
			fmt.Printf("failed to scrape: %s\n", err)
			res <- nil
		}
		if img := s.Preview.Images[0]; img != "" {
			res <- img
		} else {
			res <- nil
		}
	}()
	return res
}
