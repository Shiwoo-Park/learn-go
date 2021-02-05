package urlcheck

import (
	"errors"
	"fmt"
	"net/http"
)

var errorRequestFailed = errors.New("Request failed")

func hitURL(url string) error {
	fmt.Println("Checking URL: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode > 400 {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res.StatusCode)
		}

		return errorRequestFailed
	}
	return nil
}

func CheckV1() {
	urls := []string{
		"http://oniondev.egloos.com/9787596",
		"https://nomadcoders.co/go-for-beginners/lectures/1519",
		"https://sendbird.atlassian.net/issues/?filter=11934",
		"https://sentry.io/organizations/sendbird/issues/2178617489/?project=125361&query=is%3Aunresolved&statsPeriod=1h",
		"https://embeddedartistry.com/blog/2018/02/19/code-cleanup-splitting-up-git-commits-in-the-middle-of-a-branch/",
		"https://spin.sbtools.io/#/applications/webawsregional/executions?pipeline=gate-api-production",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.aleijfa,skdnfalejflaejf.com",
	}

	// 2 ways to initialize map
	// var results = map[string]string
	var results = make(map[string]string)

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	fmt.Println("=== Check Results ===")
	for url, result := range results {
		fmt.Println(url, ": ", result)
	}
}
