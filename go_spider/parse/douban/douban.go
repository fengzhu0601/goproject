package douban

import (
	"fengzhu0601/goproject/go_spider/fake"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DoubanMovie struct {
	Id       int
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

type Page struct {
	Page int
	Url  string
}

func ParseUrl(url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Add("Accept-Encoding", "br")
	//req.Header.Add("Accept-Languag", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Content-Length", "25")
	req.Header.Add("Content-Type", "text/html; charset=utf-8") // 按utf-8的格式返回
	//req.Header.Add("Cookie", "_ga=GA1.2.161331334.1522592243; "+
	//	"user_trace_token=20180401221723-"+uuid.GetUUID()+"; "+
	//	"LGUID=20180401221723-"+uuid.GetUUID()+"; "+
	//	"index_location_city=%E6%B7%B1%E5%9C%B3; "+
	//	"JSESSIONID="+uuid.GetUUID()+"; "+
	//	"_gid=GA1.2.1140631185.1523090450; "+
	//	"Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1522592243,1523090450; "+
	//	"TG-TRACK-CODE=index_search; _gat=1; "+
	//	"LGSID=20180407221340-"+uuid.GetUUID()+"; "+
	//	"PRE_UTM=; PRE_HOST=; PRE_SITE=https%3A%2F%2Fwww.lagou.com%2F; "+
	//	"PRE_LAND=https%3A%2F%2Fwww.lagou.com%2Fjobs%2Flist_golang%3FlabelWords%3D%26fromSearch%3Dtrue%26suginput%3D; "+
	//	"Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1523110425; "+
	//	"LGRID=20180407221344-"+uuid.GetUUID()+"; "+
	//	"SEARCH_ID="+uuid.GetUUID()+"")
	//req.Header.Add("Host", "movie.douban.com")
	//req.Header.Add("Origin", "https://movie.douban.com")
	//req.Header.Add("Referer", "https://movie.douban.com/top250")
	req.Header.Add("User-Agent", fake.GetUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	//defer resp.Body.Close()
	//fmt.Println("resp:", resp, "body:", resp.Body)
	return resp

	// 解决压缩问题
	//reader, err := gzip.NewReader(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return nil
	//}
	//fmt.Println("body:", string(body), nil)

	//////解决中文乱码
	//////bodystr := mahonia.NewDecoder("utf8").ConvertString(string(body))
	////bodystr := mahonia.NewDecoder("utf8").ConvertString(string(body))
	//fmt.Println(bodystr)
}

// 获取分页
func GetPages(url string) []Page {
	resp := ParseUrl(url)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return ParsePages(doc)
}

// 分析分页
func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

// 分析电影数据
func ParseMovies(doc *goquery.Document) (movies []DoubanMovie) {
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		id, _ := strconv.Atoi(s.Find("em").Text())

		title := s.Find(".hd a span").Eq(0).Text()

		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		movie := DoubanMovie{
			Id:       id,
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		movies = append(movies, movie)
	})

	return movies
}
