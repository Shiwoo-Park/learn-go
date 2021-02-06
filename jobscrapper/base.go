package jobscrapper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/shiwoo-park/learngo/modules"
)

type JobInfo struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL = "https://kr.indeed.com/jobs?q=%EA%B0%9C%EB%B0%9C%EC%9E%90&l=%ED%8C%90%EA%B5%90"

func cleanString(s string) string {
	txtArray := strings.Fields(strings.TrimSpace(s))
	return strings.Join(txtArray, " ")
}

func visitPage(page int, c chan *JobInfo) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)

	fmt.Println("visitPage:", pageURL)
	res, err := http.Get(pageURL)
	defer res.Body.Close()
	modules.HandlelErrorForResponse(res, err)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	modules.HandlelError(err)
	doc.Find("#pageContent .jobsearch-SerpJobCard").Each(func(i int, s *goquery.Selection) {
		attrVal, _ := s.Attr("data-jk")
		jobID := attrVal
		title := cleanString(s.Find(".title>a").Text())
		location := cleanString(s.Find(".sjcl").Text())
		salary := cleanString(s.Find(".salaryText").Text())
		summary := cleanString(s.Find(".summary").Text())
		jobInfo := JobInfo{id: jobID, title: title, location: location, salary: salary, summary: summary}
		c <- &jobInfo
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
	pageCount := 0
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html())
		pageCount = s.Find("a").Length()
	})

	var c = make(chan *JobInfo)
	for i := 0; i < pageCount; i++ {
		go visitPage(i, c)
	}

	var jobs []*JobInfo
	for i := 0; i < pageCount*10; i++ {
		jobInfo := <-c
		jobs = append(jobs, jobInfo)
	}

	for i := 0; i < len(jobs); i++ {
		fmt.Printf("%+v\n\n", jobs[i])
	}

}
