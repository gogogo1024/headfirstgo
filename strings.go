package headfirstgo

import (
	"bufio"
	"os"
)

// 读取文件，返回字符串数组
func GetString(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line :=  scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}