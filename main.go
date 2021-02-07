package main

import (
	"fmt"
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

func scrape(c echo.Context) error {
	term := strings.ToLower(modules.CleanString(c.FormValue("term")))
	jobscrapper.Scrape(term)
	return c.File("resources/jobinfos.csv")
}
