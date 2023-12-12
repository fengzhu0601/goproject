package common

import (
	"fmt"
	"io/ioutil"
	"movie_metadata/utils"
	"os"
)

// 打开文件夹，遍历文件后缀名为MP4、wmv、avi的文件，并重新命名文件
func RenameFile(dir string) {
	// 遍历文件夹
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := 0
	for _, file := range dirFiles {
		// 判断是否是文件夹
		if file.IsDir() {
			continue
		}
		// 判断文件后缀名
		if utils.IsMovieFile(file.Name()) {
			n++
			// 新文件名
			newFileName, err := utils.ExtractAndUpperCase(utils.GetFileNameWithoutExtension(file.Name()))
			if err != nil {
				fmt.Println("err :", file.Name())
				continue
			}
			newFileName += utils.GetFileExtension(file.Name())

			// 检查是否有重名文件
			_, err = os.Stat(dir + "/" + newFileName)
			if !os.IsNotExist(err) {
				fmt.Println("检查是否有重名文件:", newFileName, err)
				continue
			}

			//重命名文件
			err = os.Rename(dir+"/"+file.Name(), dir+"/"+newFileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(n, ":", "old:", file.Name(), "==> new:", newFileName)
		}
	}
}
