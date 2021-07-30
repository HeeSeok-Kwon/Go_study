package main

import (
	"Nomad_Go_Start/modules/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	// fmt.Println(c.FormValue("term"))
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName) // 첨부파일을 리턴하는 기능
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

// go get github.com/labstack/echo/v4 --> import error
// 강의에서는 error 때문에 v4 제거함
// 다시 go get github.com/labstack/echo
// file encoding error --> 한글이 깨짐
