package main

import (
	"fmt"
	"movie_metadata/utils"
	"regexp"
	"testing"
)

func TestInsertHyphen(t *testing.T) {
	input := "hello123world456"
	expectedOutput := "hello-123world-456"

	output := utils.InsertHyphen(input)

	if output != expectedOutput {
		t.Errorf("Expected %s, but got %s", expectedOutput, output)
	}

	re := regexp.MustCompile(`[a-zA-Z]+-\d+`)
	if !re.MatchString(output) {
		t.Errorf("Output doesn't match the expected pattern")
	}
	fmt.Println(output)
}

func TestIsValidDirName(t *testing.T) {
	input := "DV-1453"
	expectedOutput := true

	output := utils.IsValidDirName(input)

	if output != expectedOutput {
		t.Errorf("Expected %v, but got %v", expectedOutput, output)
	}
}

func main() {
	// 运行测试函数
	//TestInsertHyphen(nil)
	TestIsValidDirName(nil)
}
