package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fengzhu0601/goproject/go_tool/spider"
	"log"
	"net/http"
)

// 获取url的doc
func GetDoc(url string) *goquery.Document {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("创建请求失败：%v\n", err)
		log.Fatal(err)
		return nil
	}
	request.Header.Set("User-Agent", spider.GetUserAgent())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("请求网页失败：%v\n", err)
		log.Fatal(err)
		return nil
	}
	defer response.Body.Close()

	fmt.Println("url:", url, "====", response.StatusCode)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Printf("解析HTML失败：%v\n", err)
		log.Fatal(err)
		return nil
	}
	return doc
}
