/**
 * @author [gogogo1024]
 * @email [jxycbjhc@163.com]
 * @create date 2022-08-22 12:10:28
 * @modify date 2022-08-22 12:10:28
 * @desc [description] 映射
 */

package headfirstgo

import (
	"fmt"
	"log"
)

func Count() {
	lines,err := GetString("votes.txt")
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
