package utils

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// 1. 获取说有名称id
func GetNameList() []string {
	url := "https://javdb367.com/actors/6qA7?sort_type=0&t=c%2Cs"

	nameList := make([]string, 0)
	fmt.Println(nameList)

	doc := GetDoc(url)
	if doc == nil {
		return nil
	}

	doc.Find("#movie-list .item").Each(func(i int, s *goquery.Selection) {
		name := s.Find("strong").First().Text()
		fmt.Println("name", name)
	})

	return nameList
}

// 2. 遍历文件夹，找到文件相同的文件夹，进入文件夹，把文件重命名
