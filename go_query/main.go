package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var namelist = make(map[string]bool, 0)

func main() {
	delJpg()
}

// 重命名文件
func renameFile() {
	// 打开文件
	file, err := os.Open("temp.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建 CSV 读取器
	reader := csv.NewReader(file)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// 处理记录
	for _, record := range records {
		parts := strings.Split(record[0], " ")
		if len(parts) > 5 {
			namelist[parts[5]] = true
		}
	}

	var dir = "/mnt/z/sata1-13971310804/小电影/keep/松下纱荣子"
	checkSameFilenames(dir)

}

// 遍历目录下的文件夹
func checkSameFilenames(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			if namelist[file.Name()] {
				fmt.Println(file.Name())
				// processDirectory(file.Name(), dir+"/"+file.Name())
			}
		}
	}

}

func processDirectory(dirName, dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			// 如果文件名包含目录名
			if strings.Contains(file.Name(), dirName) {
				// 在相同部分后面加上 -C
				newFileName := strings.Replace(file.Name(), dirName, dirName+"-C", 1)
				fmt.Println(dirName, file.Name(), newFileName)
				fmt.Println(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, newFileName))

				// 重命名文件
				err := os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, newFileName))
				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}

func delJpg() {
	var dir = "/mnt/z/sata1-13971310804/小电影/keep/松下纱荣子"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			newFiles, err := os.ReadDir(filepath.Join(dir, file.Name()))
			if err != nil {
				fmt.Println("Error reading directory:", err)
				return
			}

			for _, newFile := range newFiles {
				errName1 := file.Name() + ".jpg"
				errName2 := file.Name() + "-C" + ".jpg"
				if newFile.Name() == errName1 || newFile.Name() == errName2 {
					err := os.Remove(filepath.Join(dir+"/"+file.Name(), newFile.Name()))
					if err != nil {
						return
					}
					fmt.Println(newFile.Name(), filepath.Join(dir+"/"+file.Name(), newFile.Name()))
				}
			}
		}
	}
}
