package urlcheck

import (
	"fmt"
	"net/http"
)

type HitResult struct {
	Url        string
	StatusCode int
	AnyError   error
}

/*
c chan<- TYPE
: This means we will ONLY SEND message through this channel in the function
*/
func hitURLV2(url string, c chan<- *HitResult) {
	fmt.Println("Checking URL: ", url)
	res, err := http.Get(url)

	var result HitResult
	if err != nil {
		result = HitResult{url, -1, err}

	} else if res.StatusCode > 400 {
		result = HitResult{url, res.StatusCode, errorRequestFailed}

	} else { // success
		result = HitResult{url, res.StatusCode, nil}
	}

	c <- &result
}

func CheckV2() {
	urls := []string{
		"http://oniondev.egloos.com/9787596",
		"https://nomadcoders.co/go-for-beginners/lectures/1519",
		"https://sendbird.atlassian.net/issues/?filter=11934",
		"https://stackoverflow.com/questidfdsfsfsdf",
		"https://sentry.io/organizations/sendbird/issues/2178617489/?project=125361&query=is%3Aunresolved&statsPeriod=1h",
		"https://embeddedartistry.com/blog/2018/02/19/code-cleanup-splitting-up-git-commits-in-the-middle-of-a-branch/",
		"https://spin.sbtools.io/#/applications/webawsregional/executions?pipeline=gate-api-production",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.aleijfa,skdnfalejflaejf.com",
	}

	// channel
	c := make(chan *HitResult)

	// 2 ways to initialize map
	// var results = map[string]string
	for _, url := range urls {
		go hitURLV2(url, c)
	}

	lenURLs := len(urls)
	for i := 0; i < lenURLs; i++ {
		fmt.Printf("%+v\n", <-c)
	}
}
