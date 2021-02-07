package jobscrapper

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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

func visitPage(url string, page int, c chan<- *JobInfo) {
	pageURL := url + "&start=" + strconv.Itoa(page*10)

	fmt.Println("visitPage:", pageURL)
	res, err := http.Get(pageURL)
	defer res.Body.Close()
	modules.HandlelErrorForResponse(res, err)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	modules.HandlelError(err)

	var extractChan = make(chan *JobInfo)
	extractCount := 0
	doc.Find("#pageContent .jobsearch-SerpJobCard").Each(func(i int, s *goquery.Selection) {
		go extractJob(s, extractChan)
		extractCount += 1
	})

	for i := 0; i < extractCount; i++ {
		c <- <-extractChan

		// above equals to logic below (receive message and send it directly)
		// extractedJob := <-extractChan
		// c <- extractedJob
	}
}

func extractJob(s *goquery.Selection, c chan *JobInfo) {
	attrVal, _ := s.Attr("data-jk")
	jobID := attrVal
	title := modules.CleanString(s.Find(".title>a").Text())
	location := modules.CleanString(s.Find(".sjcl").Text())
	salary := modules.CleanString(s.Find(".salaryText").Text())
	summary := modules.CleanString(s.Find(".summary").Text())
	jobInfo := JobInfo{id: jobID, title: title, location: location, salary: salary, summary: summary}
	c <- &jobInfo
}

func writeToCsv(jobInfoList []*JobInfo) {
	filePath := "resources/jobinfos.csv"
	f, err := os.Create(filePath)
	modules.HandlelError(err)
	defer f.Close()

	_, err = f.WriteString("link\ttitle\tlocation\tsalary\tsummary\n")
	modules.HandlelError(err)
	for _, jobInfo := range jobInfoList {
		jobInfoStr := fmt.Sprintf(
			"https://kr.indeed.com/viewjob?jk=%s\t%s\t%s\t%s\t%s\n",
			jobInfo.id, jobInfo.title, jobInfo.location, jobInfo.salary, jobInfo.summary,
		)
		_, err := f.WriteString(jobInfoStr)
		modules.HandlelError(err)
	}
}

// Scrape indeed site by a term provided
func Scrape(term string) {
	// Request the HTML page.
	var baseURL = "https://kr.indeed.com/jobs?q=" + term + "&l=%ED%8C%90%EA%B5%90"
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
		go visitPage(baseURL, i, c)
	}

	var jobs []*JobInfo
	for i := 0; i < pageCount*10; i++ {
		jobInfo := <-c
		jobs = append(jobs, jobInfo)
	}

	writeToCsv(jobs)

	fmt.Printf("Extract Jobs Done !!! (%d pages, %d jobs)", pageCount, pageCount*10)
}
