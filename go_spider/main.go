package main

import (
	"fengzhu0601/goproject/go_spider/model"
	"fengzhu0601/goproject/go_spider/target_site/cwl"
)

func main() {
	//douban.Start()
	cwl.Start(2022)

	defer model.DB.Close()
}
