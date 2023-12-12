package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var (
	BaseUrl = "https://avmoo.online/cn/search/"
)

type Movie struct {
	Title         string  `xml:"title"`
	OriginalTitle string  `xml:"originaltitle"`
	Plot          string  `xml:"plot"`
	Genre         string  `xml:"genre"`
	Director      string  `xml:"director"`
	Year          string  `xml:"year"`
	Runtime       string  `xml:"runtime"`
	Actors        []Actor `xml:"actor"`
}

type Actor struct {
	Name string `xml:"name"`
}

// 获取详情页的URL
func GetMovieUrl(dir, movieName string) (href string) {
	// 发起HTTP GET请求
	url := strings.Join([]string{BaseUrl, movieName}, "")
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求网页失败：%v\n", err)
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Printf("解析HTML失败：%v\n", err)
		log.Fatal(err)
		return
	}
	// 处理多个图片
	found := false
	doc.Find("#waterfall .item").Each(func(i int, s *goquery.Selection) {
		if found {
			return
		}
		title := s.Find("span date").First().Text()

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

	//// 有多个图片则报错
	//if doc.Find("#waterfall .item").Length() > 1 {
	//	fmt.Println("该电影有多个图片")
	//	return
	//}
	//
	//// 查找详情页
	//doc.Find("a.movie-box").Each(func(i int, s *goquery.Selection) {
	//	// 获取链接的href属性值
	//	tmpHref, exists := s.Attr("href")
	//	if exists {
	//		// 打印链接的href属性值
	//		href = tmpHref
	//		fmt.Printf("链接：%s\n", href)
	//		// 退出循环
	//		return
	//	}
	//})
	return
}

func downloadImage(dir, movieName, url string) error {
	// 发送HTTP GET请求获取图片内容
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 提取图片文件名
	filename := movieName + ".jpg"
	// 创建本地文件
	file, err := os.Create(dir + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将图片内容写入本地文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("已下载图片：%s\n", filename)
	return nil
}

//// 截取图片
//func cutImage(dir, movieName string) {
//	inputPath := dir + movieName + "-poster" + ".jpg"
//	// 打开原始图片文件
//	file, err := os.Open(inputPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	// 解码原始图片
//	img, _, err := image.Decode(file)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 获取原始图片的尺寸
//	bounds := img.Bounds()
//
//	// 计算左上角和右下角坐标
//	x1 := bounds.Max.X - 410
//	y1 := bounds.Min.Y
//	x2 := bounds.Max.X
//	y2 := bounds.Max.Y
//
//	// 裁剪图片
//	croppedImg := img.(interface {
//		SubImage(r image.Rectangle) image.Image
//	}).SubImage(image.Rect(x1, y1, x2, y2))
//
//	// 保存裁剪后的图片
//	outputPath := dir + movieName + ".jpg"
//
//	// 创建输出文件
//	outputFile, err := os.Create(outputPath)
//	if err != nil {
//		log.Fatal("1111:", err)
//	}
//	defer outputFile.Close()
//
//	// 将裁剪后的图片保存为JPEG格式
//	err = jpeg.Encode(outputFile, croppedImg, &jpeg.Options{Quality: 100})
//	if err != nil {
//		log.Fatal("222:", err)
//	}
//	log.Println("图片裁剪完成:", outputPath)
//}

func GetMovieInfo(dir, movieName, href string) {
	// 发起HTTP GET请求
	url := strings.Join([]string{"https:", href}, "")
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

	// 生成nfo文件
	movie := &Movie{
		Actors: make([]Actor, 0),
	}
	movie.Title = doc.Find("h3").Text()
	movie.OriginalTitle = movieName

	findActor(doc, movie)
	//actor := Actor{
	//	Name: doc.Find("#avatar-waterfall span").Text(),
	//}
	//if actor.Name == "" {
	//	//re := regexp.MustCompile("[\\p{Han}]+$") // 匹配最后的汉字字符
	//	re := regexp.MustCompile(`[^a-zA-Z]+$`)
	//	actor.Name = re.FindString(DIR)
	//}
	//movie.Actors = append(movie.Actors, actor)

	movie.Year = extractValue(doc, "span.header", "发行时间:")
	if len(movie.Year) >= 4 {
		movie.Year = movie.Year[:4]
	}
	movie.Runtime = extractValue(doc, "span.header", "长度:")
	if len(movie.Runtime) >= 6 {
		movie.Runtime = movie.Runtime[:len(movie.Runtime)-6]
	}
	movie.Director = extractValue(doc, "span.header", "导演:")
	fmt.Println(movie)
	GenerateNfo(dir, movieName, movie)

	// 查找大图
	doc.Find("a.bigImage").Each(func(i int, s *goquery.Selection) {
		// 获取链接的href属性值
		href, exists := s.Attr("href")
		if exists {
			// 打印链接的href属性值
			fmt.Printf("链接：%s\n", href)
			// 下载图片
			downloadImage(dir, movieName+"-poster", href)
			// 图片裁剪
			cutImage(dir, movieName)
			return
		}
	})
}

// 查找演员名
func findActor(doc *goquery.Document, movie *Movie) {
	doc.Find("#avatar-waterfall .avatar-box").Each(func(i int, s *goquery.Selection) {
		actorName := s.Find("span").Text()
		fmt.Println("findActor:", actorName)
		if actorName != "" {
			movie.Actors = append(movie.Actors, Actor{Name: actorName})
		}
	})
}

func cutImage(dir, movieName string) {
	inputPath := dir + movieName + "-poster" + ".jpg"
	// 打开原始图片文件
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("444:", err)
	}
	defer file.Close()
	// 打开原始图片文件
	img, err := imaging.Open(inputPath)
	if err != nil {
		return
		//log.Fatal("5555:", err)
	}

	// 获取原始图片的尺寸
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// 计算裁剪区域的坐标
	x1 := width - 390
	y1 := 0
	x2 := width
	y2 := height

	// 裁剪图片
	croppedImg := imaging.Crop(img, image.Rect(x1, y1, x2, y2))

	// 调整裁剪后的图片尺寸
	targetWidth := 410
	resizedImg := imaging.Resize(croppedImg, targetWidth, 0, imaging.Lanczos)

	// 保存裁剪后的图片
	outputPath := dir + movieName + ".jpg"
	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal("1111:", err)
	}
	defer outputFile.Close()

	// 保存裁剪后的图片
	err = imaging.Save(resizedImg, outputPath)
	if err != nil {
		log.Fatal("3333:", err)
	}
	log.Println("图片裁剪完成:", outputPath)
}

//// 生成NFO文件
//func GenerateNfo(dir, movieName string, movie *Movie) {
//	// 创建nfo文件
//	nfoFile, err := os.Create(dir + movieName + ".nfo")
//	if err != nil {
//		fmt.Printf("创建nfo文件失败：%v\n", err)
//		log.Fatal(err)
//	}
//	defer nfoFile.Close()
//
//	// 创建一个字符串构建器
//	builder := strings.Builder{}
//	// 写入数据到构建器
//	builder.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\" ?>\n")
//	builder.WriteString("<movie>\n")
//	builder.WriteString("\t<title>" + movie.Title + "</title>\n")
//	builder.WriteString("\t<originaltitle>" + movieName + "</originaltitle>\n")
//	builder.WriteString("\t<plot>" + movie.Plot + "</plot>\n")
//	builder.WriteString("\t<genre>" + movie.Genre + "</genre>\n")
//	builder.WriteString("\t<country>日本</country>\n")
//	builder.WriteString("\t<director>" + movie.Director + "</director>\n")
//	builder.WriteString("\t<year>" + movie.Year + "</year>\n")
//	builder.WriteString("\t<runtime>" + movie.Runtime + "</runtime>\n")
//	builder.WriteString("\t<actor>\n")
//	builder.WriteString("\t\t<name>" + "愛乃なみ" + "</name>\n")
//	builder.WriteString("\t</actor>\n")
//	builder.WriteString("</movie>\n")
//
//	// 将字符串写入文件
//	_, err = nfoFile.WriteString(builder.String())
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("数据已写入 example.nfo")
//}

func extractValue(doc *goquery.Document, headerSelector, target string) string {
	var result string
	doc.Find(headerSelector).Each(func(_ int, s *goquery.Selection) {
		if strings.Contains(s.Text(), target) {
			value := s.Parent().Text()
			result = strings.TrimSpace(strings.TrimPrefix(value, target))
		}
	})
	return result
}

// 生成NFO文件
func GenerateNfo(dir, movieName string, movie *Movie) {
	// 创建nfo文件
	nfoFile, err := os.Create(dir + movieName + ".nfo")
	if err != nil {
		fmt.Printf("创建nfo文件失败：%v\n", err)
		log.Fatal(err)
	}
	defer nfoFile.Close()

	// 从文件中读取 NFO 模板
	nfoTemplateBytes, err := ioutil.ReadFile("./template/nfo.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	nfoTemplate := string(nfoTemplateBytes)

	// 解析模板字符串
	tmpl, err := template.New("nfoTemplate").Parse(nfoTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 根据模板和数据生成 NFO 内容
	err = tmpl.Execute(nfoFile, movie)
	if err != nil {
		fmt.Println(err)
		return
	}
}
