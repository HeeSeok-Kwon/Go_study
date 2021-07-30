package main

import (
	"Nomad_Go_Start/modules/scrapper"
	"strings"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	// fmt.Println(c.FormValue("term"))
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	return nil
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
