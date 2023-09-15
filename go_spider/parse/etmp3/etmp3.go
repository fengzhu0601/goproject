package etmp3

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

var (
	BaseUrl = "http://www.etmp3.com"
)

type Page struct {
	Url string
}

// 获取分页
func GetPages(url string) ([]Page, int) {
	return ParsePages(url)
}

func GetPageDoc(url string) *goquery.Document {
	// 发起HTTP GET请求
	url = strings.Join([]string{BaseUrl, url}, "")
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求网页失败：%v\n", err)
		log.Fatal(err)
	}
	defer response.Body.Close()

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Printf("解析HTML失败：%v\n", err)
		log.Fatal(err)
	}
	return doc
}

// 分析分页
func ParsePages(url string) (pages []Page, count int) {
	// 第一页
	currentPageLink := url
	fmt.Println("第一页链接:", currentPageLink)
	pages = append(pages, Page{Url: currentPageLink})

	for {
		doc := GetPageDoc(currentPageLink)
		// 提取当前页面的下一页链接
		nextPageLink, exists := doc.Find(".page a:contains('下一页')").Attr("href")

		doc.Find(".play_list ul li").Each(func(i int, s *goquery.Selection) {
			count += 1
		})
		if !exists {
			return
		}
		// 生成完整的下一页链接
		currentPageLink = nextPageLink
		fmt.Println("下一页链接:", currentPageLink)
		pages = append(pages, Page{
			Url: nextPageLink,
		})
	}
}
