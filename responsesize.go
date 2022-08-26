// 指针
package headfirstgo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL string
	Size int
}

// 获取url页面response大小，并发送给channel
func ResponseSize(url string, channel chan Page) {
	fmt.Println("getting", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL:url,Size:len(body)}

}
