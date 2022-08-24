package headfirstgo

import (
	"fmt"
	"log"
)

// 映射统计
func Count(path string) {
	lines,err := GetString(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	ranks := make(map[string]int)
	
	for _,line := range lines {
		ranks[line]++
	}
	fmt.Printf("%#v\n",ranks)

	for key,value := range ranks{
		fmt.Printf("%s:%d\n",key,value)
	}
}
