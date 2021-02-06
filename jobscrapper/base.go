package jobscrapper

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/shiwoo-park/learngo/modules"
)

var baseURL = "https://kr.indeed.com/jobs?q=%EA%B0%9C%EB%B0%9C%EC%9E%90&l=%ED%8C%90%EA%B5%90"

func visitPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page)

	res, err := http.Get(pageURL)
	defer res.Body.Close()
	modules.HandlelErrorForResponse(res, err)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	modules.HandlelError(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html())
	})
}

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get(baseURL)
	defer res.Body.Close()
	modules.HandlelErrorForResponse(res, err)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	modules.HandlelError(err)

	// Get total page number
	pages := 0
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html())
		pages = s.Find("a").Length()
	})

	for 

}
