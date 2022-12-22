package main

import (
	"fengzhu0601/goproject/go_spider/model"
	"fengzhu0601/goproject/go_spider/parse"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

// 新增数据
func Add(movies []parse.DoubanMovie) {
	for index, movie := range movies {
		if err := model.DB.Create(&movie).Error; err != nil {
			log.Printf("db.Create index: %s, err : %v", index, err)
		}
	}
}

// 开始爬取
func Start() {
	// 创建数据库表
	model.DB.AutoMigrate(&parse.DoubanMovie{})

	var movies []parse.DoubanMovie

	pages := parse.GetPages(BaseUrl)
	for _, page := range pages {
		resp := parse.ParseUrl(strings.Join([]string{BaseUrl, page.Url}, ""))
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		movies = append(movies, parse.ParseMovies(doc)...)
	}
	Add(movies)
}

func main() {
	Start()

	defer model.DB.Close()
}
