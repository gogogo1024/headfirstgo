package headfirstgo

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 打开文件
func OpenFile(filename string)(*os.File,error)  {
	fmt.Println("opening file: ",filename)
	return  os.Open(filename)
}
// 关闭文件
func CloseFile(file *os.File)  {
	fmt.Println("closing file")
	file.Close()
}
// 读取文件
func ReadFile(path string) {
	fmt.Println("read file start")
	file, err :=OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// 保证文件在后面代码执行完之后，在退出ReadFile之前被关闭
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())

	}
	fmt.Println("read file end")
}
