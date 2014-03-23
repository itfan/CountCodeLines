package main

import (
	"countlines"
	"fmt"
	"log"
	"os"
)

func main() {
	startPath, err := os.Getwd() //与src目录同级
	if err != nil {
		log.Fatal(err)
	}

	mcountlines := countlines.Countlines{}
	mcountlines.StartPath = startPath
	fmt.Println(mcountlines.TotalLines())

	var i int
	_, err = fmt.Scanf("%d", &i)
}
