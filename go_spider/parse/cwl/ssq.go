package cwl

import (
	"fengzhu0601/goproject/go_spider/fake"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type CwlSsq struct {
	Issue     string // 期号
	Red1      string //
	Red2      string //
	Red3      string //
	Red4      string //
	Red5      string //
	Red6      string //
	Blue1     string
	HappyWeek string
	Jackpot   string
	No1       string
	No1Money  string
	No2       string
	No2Money  string
	BetAmount string // 总投注额
	AwardDate string // 开奖日期
}

var YearUrl = "http://datachart.500.com/ssq/history/newinc/history.php?start=%s&end=%s"

// 爬取某年的所有数据
func GetYear(year int32) []CwlSsq {
	url := fmt.Sprintf(YearUrl, "01001", "23060")
	fmt.Println(url)
	resp := ParseUrl(url)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return ParseSSQ(doc)
}

func ParseUrl(url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Add("Accept-Encoding", "gzip, deflate")
	//req.Header.Add("Accept-Languag", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	//req.Header.Add("Connection", "keep-alive")
	////req.Header.Add("Content-Length", "25")
	//req.Header.Add("Content-Type", "text/json;charset=UTF-8") // 按utf-8的格式返回
	//req.Header.Add("User-Agent", fake.GetUserAgent())
	req.Header.Set("User-Agent", fake.GetUserAgent())
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	//defer resp.Body.Close()
	fmt.Println("resp:", resp, "body:", resp.Body)
	return resp
}

// 分析双色球数据
func ParseSSQ(doc *goquery.Document) (ssqs []CwlSsq) {
	doc.Find("#tdata > tr.t_tr1").Each(func(i int, s *goquery.Selection) {
		var tdList []string
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			tdList = append(tdList, s.Text())
		})
		ssq := CwlSsq{
			Issue:     tdList[0],
			Red1:      tdList[1],
			Red2:      tdList[2],
			Red3:      tdList[3],
			Red4:      tdList[4],
			Red5:      tdList[5],
			Red6:      tdList[6],
			Blue1:     tdList[7],
			HappyWeek: tdList[8],
			Jackpot:   tdList[9],
			No1:       tdList[10],
			No1Money:  tdList[11],
			No2:       tdList[12],
			No2Money:  tdList[13],
			BetAmount: tdList[14],
			AwardDate: tdList[15],
		}
		fmt.Println(ssq)
		ssqs = append(ssqs, ssq)
	})

	return ssqs
}
