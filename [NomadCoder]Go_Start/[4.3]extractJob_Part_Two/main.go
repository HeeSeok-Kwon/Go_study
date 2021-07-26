package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	var jobs []extractedJob
	// pages := getPages()
	totalPages := getPages()
	// fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...) // contents [] [] [] [] [] ... -> [] 하나로 합치는 역할
	}
	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 이 함수 끝났을 때 닫아줘야 함
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCard := doc.Find(".tapItem")
	searchCard.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	// id, _ := card.Attr("href")
	// fmt.Println(id)
	title := cleanString(card.Find(".jobTitle>span").Text())
	// fmt.Println(title)
	location := cleanString(card.Find(".companyLocation").Text())
	// fmt.Println(location)
	salary := cleanString(card.Find(".salary-snippet").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	// fmt.Println(id, title, location, salary, summary)
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 이 함수 끝났을 때 닫아줘야 함
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// For next lecture -- Each()
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html())
		// fmt.Println(s.Find("a").Length())
		pages = s.Find("a").Length()
	})
	// fmt.Println(doc)
	return pages
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
