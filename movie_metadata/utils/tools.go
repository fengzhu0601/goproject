package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	//DIR = "/mnt/z/sata1-13971310804/小电影/Japan/julia"
	DIR = "/mnt/z/sata1-13971310804/小电影/Japan/松下纱荣子"
)

func ExtractAndUpperCase(input string) (string, error) {
	re := regexp.MustCompile(`[a-zA-Z]+-\d+`)
	result := re.FindString(input)
	if result == "" {
		re = regexp.MustCompile(`[a-zA-Z]+-\d+|[a-zA-Z]+-\d+[a-zA-Z]+-\d+|[a-zA-Z]+\d+`)
		result = re.FindString(input)
	}
	result = strings.ToUpper(result)

	if result == "" {
		return "", errors.New("无法提取有效字符串")
	}

	return InsertHyphen(result), nil
}

func InsertHyphen(input string) string {
	re := regexp.MustCompile(`([a-zA-Z]+)(\d+)`)
	return re.ReplaceAllString(input, "$1-$2")
}

// 获取文件的后缀名
func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}

// 获取不包括后缀名的文件名
func GetFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, GetFileExtension(fileName))
}

// 判断目录下是否有.jpg格式的文件
func HasJPEGFile(dir, format string) bool {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, file := range dirFiles {
		// 判断是否是文件夹
		if file.IsDir() {
			continue
		}
		if GetFileExtension(file.Name()) == format {
			return true
		}
	}
	return false
}

func MoveFilesInDirectory(directoryPath string) {
	// 遍历目录
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			dirPath := filepath.Join(directoryPath, file.Name())

			// 判断目录名是否符合规则
			match, _ := regexp.MatchString("^[a-zA-Z]+-[0-9]+$", file.Name())
			if !match {
				// 移动目录中的文件到父目录
				dirFiles, _ := ioutil.ReadDir(dirPath)
				for _, dirFile := range dirFiles {
					filePath := filepath.Join(dirPath, dirFile.Name())
					newPath := filepath.Join(directoryPath, dirFile.Name())
					err := os.Rename(filePath, newPath)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
				// 删除目录
				err := os.Remove(dirPath)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				// 递归处理子目录
				MoveFilesInDirectory(dirPath)
			}
		}
	}
}

// 判断目录名是否符合规则
func IsValidDirName(dirName string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]+-[0-9]+$", dirName)
	return match
}
