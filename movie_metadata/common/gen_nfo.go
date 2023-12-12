package common

import (
	"fmt"
	"io/ioutil"
	"movie_metadata/utils"
)

// 获取目录下每个文件夹的图片信息
func GenMoviesNfo(dir string) {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range dirFiles {
		// 判断是否是文件夹
		if file.IsDir() {
			if utils.IsValidDirName(file.Name()) {
				//time.Sleep(10 * time.Second)
				// 获取文件夹下图片信息
				href := utils.GetMovieUrl(file.Name())
				if href == "" {
					fmt.Println("GEnMOviesNFO:", file.Name())
					continue
				}
				utils.GetMovieInfo(dir+"/"+file.Name()+"/", file.Name(), href)
			} else {
				GenMoviesNfo(dir + "/" + file.Name())
			}
		}
	}
}
