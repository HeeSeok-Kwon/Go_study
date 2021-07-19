package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	// pages := getPages()
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 이 함수 끝났을 때 닫아줘야 함
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// For next lecture -- Each()
	doc.Find("./pagination").Each()
	fmt.Println(doc)
	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
// terminal --> go get github.com/PuerkitoBio/goquery 
