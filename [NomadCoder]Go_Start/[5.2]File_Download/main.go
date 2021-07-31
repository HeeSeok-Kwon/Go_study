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
// 해결방안
// (1) 다운받은 csv file 우클릭 --> 메모장으로 열기
// (2) 엑셀 빈 파일을 열고 --> 데이터 --> 새쿼리 --> 파일에서 --> csv에서 --> jobs.csv 선택 --> 가져오기 --> 파일원본 드롭다운 버튼 클릭 후 65001:유니코드(UTF-8) 선택 --> 로드 클릭
// --> 디자인 탭 --> 머리글 행 체크 해제 --> 첫 번째 행 삭제 --> 표스타일 : 없음 선택 --> 저장
