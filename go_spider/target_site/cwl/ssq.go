package cwl

import (
	"fengzhu0601/goproject/go_spider/model"
	"fengzhu0601/goproject/go_spider/parse/cwl"
	"fmt"
	"log"
)

var (
	BaseUrl = "http://www.cwl.gov.cn/ygkj/wqkjgg/ssq/"
)

// 新增数据
func Add(ssqs []cwl.CwlSsq) {
	for index, ssq := range ssqs {
		fmt.Println(index, ssq)
		if err := model.DB.Create(&ssq).Error; err != nil {
			log.Printf("db.Create index: %s, err : %v", index, err)
		}
	}
}
func Start(year int32) {
	//创建数据库表
	model.DB.AutoMigrate(&cwl.CwlSsq{})

	//cwl.GetYear(year)
	ssqs := cwl.GetYear(year)
	fmt.Println(len(ssqs))
	Add(ssqs)
}
