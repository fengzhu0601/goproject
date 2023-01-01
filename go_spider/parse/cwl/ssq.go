package cwl

import (
	"fengzhu0601/goproject/go_spider/fake"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SsqResult struct {
	State     int32
	Message   string
	PageCount int32
	CountNum  int32
	Tflag     int32
	Result    []CwlSsq
}

type CwlSsq struct {
	Name        string `json:"name"`
	Code        string
	DetailsLink string
	VideoLink   string
	Date        string
	Week        string
	Red         string
	Blue        string
	Blue2       string
	Sales       string
	PoolMoney   string
	Content     string
	AddMoney    string
	AddMoney2   string
	Msg         string
	Z2Add       string
	Z2Add2      string
	PrizeGrades []PrizeGrade
}
type PrizeGrade struct {
	Type      int32
	TypeNum   string
	TypeMoney string
}

var YearUrl = "http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=%s&issueEnd=%s&dayStart=&dayEnd="

// 爬取某年的所有数据
func GetYear(year int32) []CwlSsq {
	url := fmt.Sprintf(YearUrl, "2022002", "2022100")
	fmt.Println(url)
	resp := ParseUrl(url)
	fmt.Println(resp)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
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
	req.Header.Set("Host", "www.cwl.gov.cn")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	//defer resp.Body.Close()
	fmt.Println("resp:", resp, "body:", resp.Body)
	return resp
}
