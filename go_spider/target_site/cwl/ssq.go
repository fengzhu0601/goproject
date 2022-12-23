package cwl

import (
	"fengzhu0601/goproject/go_spider/parse/cwl"
)

var (
	BaseUrl = "http://www.cwl.gov.cn/ygkj/wqkjgg/ssq/"
)

func Start(year int32) {
	cwl.GetYear(year)
}
