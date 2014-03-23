package countlines

import (
	"bufio"
	//"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Countlines struct {
	StartPath string
	total     int
}

func (this *Countlines) TotalLines() int {
	filepath.Walk(this.StartPath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			//println(path)
			num := GoFileLines(path)
			this.total += num
			//fmt.Println(num)
			return nil
		})

	return this.total
}

//给定一个文件输出它有多少行
func GoFileLines(filePath string) int {
	var lineNum int
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	mbuf := bufio.NewReader(file)

	for {
		_, err := mbuf.ReadString('\n') //以'\n'为结束符读入一行
		if true {
			lineNum++
		}
		if err != nil || io.EOF == err {
			break
		}

	}

	file.Close()
	return lineNum
}
