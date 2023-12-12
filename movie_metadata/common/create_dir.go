package common

import (
	"fmt"
	"io/ioutil"
	"movie_metadata/utils"
	"os"
	"regexp"
)

// 为每个文件创建一个文件夹
func CreateDir(dir string) {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range dirFiles {
		// 判断是否是文件夹
		if file.IsDir() {
			continue
		}
		// 判断文件名是否符合规则
		re := regexp.MustCompile(`^[a-zA-Z]+-[0-9]+$`)
		if !re.MatchString(utils.GetFileNameWithoutExtension(file.Name())) {
			continue
		}

		// 判断文件后缀名
		if utils.IsMovieFile(file.Name()) {
			// 创建文件夹
			err := os.Mkdir(dir+"/"+utils.GetFileNameWithoutExtension(file.Name()), os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
			// 移动文件
			err = os.Rename(dir+"/"+file.Name(), dir+"/"+utils.GetFileNameWithoutExtension(file.Name())+"/"+file.Name())
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
