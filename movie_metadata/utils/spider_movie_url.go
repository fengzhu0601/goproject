package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// 获取详情页的URL
func GetMovieUrl(movieName string) (href string) {
	// 发起HTTP GET请求
	url := strings.Join([]string{BaseUrl, movieName}, "")
	doc := GetDoc(url)
	if doc == nil {
		return
	}
	fmt.Println(doc)

	// 处理多个图片,取第一个
	found := false
	doc.Find("#waterfall .item").Each(func(i int, s *goquery.Selection) {
		if found {
			fmt.Println("found:", movieName)
			return
		}
		title := s.Find("span date").First().Text()
		fmt.Println("title:", title, movieName)

		title = strings.ToUpper(title)
		if title == strings.ToUpper(movieName) {
			fmt.Println("title:", title, movieName)
			tmpHref, exists := s.Find("a").Attr("href")
			if exists {
				// 打印链接的href属性值
				href = tmpHref
				fmt.Printf("链接：%s\n", href)
				found = true
			}
		}
	})
	return
}

// 容易被403
//func GetMovieUrl(movieName string) (href string) {
//	// 发起HTTP GET请求
//	url := strings.Join([]string{BaseUrl, movieName}, "")
//	response, err := http.Get(url)
//	if err != nil {
//		fmt.Printf("请求网页失败：%v\n", err)
//		log.Fatal(err)
//		return
//	}
//	defer response.Body.Close()
//
//	// 解析HTML
//	doc, err := goquery.NewDocumentFromReader(response.Body)
//	if err != nil {
//		fmt.Printf("解析HTML失败：%v\n", err)
//		log.Fatal(err)
//		return
//	}
//	// 处理多个图片,取第一个
//	found := false
//	doc.Find("#waterfall .item").Each(func(i int, s *goquery.Selection) {
//		if found {
//			fmt.Println("found:", movieName)
//			return
//		}
//		title := s.Find("span date").First().Text()
//		fmt.Println("title:", title, movieName)
//
//		title = strings.ToUpper(title)
//		if title == strings.ToUpper(movieName) {
//			fmt.Println("title:", title, movieName)
//			tmpHref, exists := s.Find("a").Attr("href")
//			if exists {
//				// 打印链接的href属性值
//				href = tmpHref
//				fmt.Printf("链接：%s\n", href)
//				found = true
//			}
//		}
//	})
//	return
//}
