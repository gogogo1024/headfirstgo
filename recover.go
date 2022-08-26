package headfirstgo

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// 扫描文件夹
func ScanDir(path string)  {
	fmt.Println(path)
	files, err:=ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _,file := range files{
		filePath:=filepath.Join(path,file.Name())
		if file.IsDir() {
			ScanDir(filePath)
		}else {
			fmt.Println(filePath)
		}
	}
}

// 利用recover恢复panic
func ReportPanic(){
	p:=recover()
	err,ok:= p.(error)
	if ok {
		fmt.Println(err)
	}else {
		panic(p)
	}
}