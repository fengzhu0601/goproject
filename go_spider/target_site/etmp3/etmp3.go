package etmp3

import (
	"fengzhu0601/goproject/go_spider/parse/etmp3"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gosuri/uiprogress"
	"io"
	"net/http"
	"os"
	"strings"
)

func Start(url string) {
	saveDir := "./songs"

	// 创建保存歌曲的文件夹
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件夹失败：%v\n", err)
		return
	}

	pages, count := etmp3.GetPages(url)
	fmt.Println(count)
	uiprogress.Start()
	bar := uiprogress.AddBar(count).AppendCompleted().PrependElapsed()

	for _, page := range pages {
		doc := etmp3.GetPageDoc(page.Url)
		// 查找所有歌曲链接
		doc.Find(".play_list ul li").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Find(".name a").Attr("href")
			if exists {
				songName := s.Find(".name a").Text()
				ParseSong(link, saveDir, songName)
				bar.Incr()
			}
		})
	}
	uiprogress.Stop()
	fmt.Println("爬虫程序执行完毕:", count)
}

// 解析每首歌曲的链接
func ParseSong(link, saveDir, songName string) {
	if strings.Contains(songName, "[MP3_LRC]") {
		songName = strings.Replace(songName, "[MP3_LRC]", "", -1)
	}

	doc := etmp3.GetPageDoc(link)
	mp3link, exists := doc.Find(".dance_wl a").Attr("href")
	if exists {
		// 下载歌曲
		downloadSong(mp3link, saveDir, songName)
	} else {
		fmt.Println("找不到", songName)
	}

	//lrcLink, exists := doc.Find(".dance_wl a:contains('LRC')").Attr("href")
	//if exists {
	//	// 下载歌词
	//	downloadLRC(lrcLink, saveDir, songName)
	//} else {
	//	fmt.Println("找不到LRC歌词", songName)
	//}
}

// 下载歌曲
func downloadSong(url, saveDir, songName string) {
	fileName := songName + ".mp3"
	downloadFile(url, saveDir, fileName)
}

// 下载LRC歌词
func downloadLRC(url, saveDir, songName string) {
	fileName := songName + ".lrc"
	downloadFile(url, saveDir, fileName)
}

// 下载文件
func downloadFile(url, saveDir, name string) {
	url = strings.Join([]string{etmp3.BaseUrl, url}, "")
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("下载歌曲失败：%v\n", err)
		return
	}
	defer response.Body.Close()

	// 提取文件名
	//tokens := strings.Split(url, "/")
	//fileName := tokens[len(tokens)-1]

	// 创建保存文件
	file, err := os.Create(saveDir + "/" + name)
	if err != nil {
		fmt.Printf("创建文件失败：%v\n", err)
		return
	}
	defer file.Close()

	// 将歌曲内容写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("保存文件失败：%v\n", err)
		return
	}

	fmt.Printf("歌曲 %s 下载完成\n", name)
}
