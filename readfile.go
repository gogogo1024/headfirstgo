package headfirstgo

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(path string) {
	fmt.Println("read file start")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())

	}
	fmt.Println("read file end")
}
