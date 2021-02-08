package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shiwoo-park/learngo/jobscrapper"
	"github.com/shiwoo-park/learngo/modules"
)

type BigStruct struct {
	Data [10]byte
}

func main() {
	fmt.Println("START LEAN GO !!!!!")

	// #1 ---- Theory ----
	// modules.LearnBasics()
	// modules.StrFormat()

	// #2 ---- Bank & Dict Project ----
	// banking.BankingProject()
	// modules.DicTypeTest()

	// #3 ---- URL Checker & Go Routine ----
	// urlcheck.CheckV1()
	// modules.ExperienceChannelAndGoRoutine()
	// urlcheck.CheckV2()

	// #4 ---- Job Scrapper ----
	// term := "python"
	// jobscrapper.Scrape(term)

	// Echo instance
	echoServer := echo.New()

	// Middleware
	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	// Routes
	echoServer.GET("/", home)
	echoServer.POST("/scrape", scrape)

	// Start server
	echoServer.Logger.Fatal(echoServer.Start(":1323"))

}

func home(c echo.Context) error {
	return c.File("static/home.html")
	// return c.String(http.StatusOK, "hello world !!!")
}

// scrape by term and download csv file
func scrape(c echo.Context) error {
	defer os.Remove("resources/jobinfos.csv")

	term := strings.ToLower(modules.CleanString(c.FormValue("term")))
	jobscrapper.Scrape(term)

	// both act in the same way
	return c.Attachment("resources/jobinfos.csv", "jobinfos.csv")
	// return c.File("resources/jobinfos.csv")
}
