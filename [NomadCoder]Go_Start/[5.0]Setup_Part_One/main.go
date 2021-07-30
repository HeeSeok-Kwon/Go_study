package main

import "Nomad_Go_Start/modules/scrapper"

func main() {
	scrapper.Scrape("term")
}

// go get github.com/labstack/echo/v4
// 강의에서는 error 때문에 v4 제거함
