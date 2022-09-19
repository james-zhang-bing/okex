package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {
	ps := scanFile("/home/james/go/src/github.com/okex/models")
	for _, v := range ps {
		fmt.Println(v)
	}

	j, err := ioutil.ReadFile("./comment2.json")
	if err != nil {
		log.Fatal(err)
	}
	var m map[string][]string
	err = json.Unmarshal(j, &m)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range ps {
		rewirte(f, m)
	}

}

func scanFile(pw string) []string {
	var paths []string
	dir, err := os.ReadDir(pw)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dir {
		if d.IsDir() {
			gofile := scanFile(path.Join(pw, d.Name()))
			paths = append(paths, gofile...)
			continue
		}
		paths = append(paths, path.Join(pw, d.Name()))
	}
	return paths
}

func rewirte(filePath string, m map[string][]string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	regtag := regexp.MustCompile("`json:\"(\\w+)")
	//regtag := regexp.MustCompile("`json:\"\\w+[\\S]+\"`")
	var newFileStr []string
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Println("file end")
				break
			}
			log.Fatal(err)
		}
		//fmt.Println(string(l))
		j := regtag.FindStringSubmatch(string(l))
		//fmt.Printf("%#v\n", j)

		if len(j) >= 2 {
			if com, ok := m[j[1]]; ok {
				c := strings.Count(string(l), "\t")
				pre := strings.Repeat("\t", c)
				newFileStr = append(newFileStr, fmt.Sprintf("%s// %s", pre, strings.Join(com, " ")))
			}

		}
		newFileStr = append(newFileStr, string(l))
		
	}
	f.Close()
	os.Remove(filePath)
	f, err = os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, v := range newFileStr {
		fmt.Println(v)
		_,err:=w.WriteString(fmt.Sprintf("%s\n", v))
		if err!=nil{
			log.Fatal(err)
		}
	}
	w.Flush()
	//f.Sync()
}
