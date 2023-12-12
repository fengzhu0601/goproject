package main

import (
	"fmt"
	"io/ioutil"
	"movie_metadata/utils"
	"os"
	"regexp"
)

func main() {

	//renameFile(utils.DIR)
	//createDir(utils.DIR)
	getDirImgInfo(utils.DIR)

	//utils.MoveFilesInDirectory(utils.DIR)

	//utils.GetMovieInfo(utils.DIR+"/"+"STARS-236"+"/", "STARS-236", "//avmoo.online/cn/movie/7194ccaf2da7387f")
}

// 打开文件夹，遍历文件后缀名为MP4、wmv、avi的文件，并重新命名文件
func renameFile(dir string) {
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
		if utils.GetFileExtension(file.Name()) == ".mp4" ||
			utils.GetFileExtension(file.Name()) == ".wmv" ||
			utils.GetFileExtension(file.Name()) == ".MP4" ||
			utils.GetFileExtension(file.Name()) == ".rmvb" ||
			utils.GetFileExtension(file.Name()) == ".avi" ||
			utils.GetFileExtension(file.Name()) == ".mkv" {
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

// 为每个文件创建一个文件夹
func createDir(dir string) {
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
		if utils.GetFileExtension(file.Name()) == ".mp4" ||
			utils.GetFileExtension(file.Name()) == ".wmv" ||
			utils.GetFileExtension(file.Name()) == ".MP4" ||
			utils.GetFileExtension(file.Name()) == ".rmvb" ||
			utils.GetFileExtension(file.Name()) == ".avi" ||
			utils.GetFileExtension(file.Name()) == ".AVI" ||
			utils.GetFileExtension(file.Name()) == ".mkv" {
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

// 获取目录下每个文件夹的图片信息
func getDirImgInfo(dir string) {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range dirFiles {
		// 判断是否是文件夹
		if file.IsDir() {
			////如果已经下载图片了就不要再下载了
			//if utils.HasJPEGFile(dir + "/" + file.Name() + "/") {
			//	continue
			//}
			//fmt.Println(dir + "/" + file.Name() + "/")
			// 获取文件夹下图片信息
			href := utils.GetMovieUrl(dir+"/"+file.Name()+"/", file.Name())
			if href == "" {
				continue
			}
			utils.GetMovieInfo(dir+"/"+file.Name()+"/", file.Name(), href)
		}
	}
}
