package expand

import (
	"net/http"
)

func Get(url string) (string, error) {
	client := &http.Client{}
	resp, err := client.Head(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// redirect
	if location, err := resp.Location(); err == nil && location != nil && resp.StatusCode/100 == 3 {
		return Get(location.String())
	}

	if resp.Request != nil {
		return resp.Request.URL.String(), nil
	}

	return url, nil
}

func Gets(urls []string) []string {
	expandUrls := make([]string, 0)
	for _, url := range urls {
		url, err := Get(url)
		if err == nil {
			expandUrls = append(expandUrls, url)
		}
	}

	return expandUrls
}
