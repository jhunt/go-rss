package rss

import (
	"net/http"
)

func Fetch(url string) (Feed, error) {
	res, err := http.Get(url)
	if err != nil {
		return Feed{}, err
	}
	return ReadClose(res.Body)
}
