package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func genJSON() {
	f, err := os.Open("./okapi-zh.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	//> result
	rex := regexp.MustCompile(`\w+`)

	m := make(map[string][]string)
	table := doc.Find("table")
	table.Each(func(i int, s *goquery.Selection) {
		s.Find("tbody").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				trs := make([]string, 0)
				s.Find("td").Each(func(i int, s *goquery.Selection) {
					trs = append(trs, Text(s))
				})
				if len(trs) == 0 {
					return
				}
				for i, v := range trs {
					l := strings.TrimLeft(v, " ")
					l = strings.TrimRight(l, " ")
					trs[i] = l
				}

				key := rex.FindString(trs[0])
				if key == "" {
					return
				}
				//fmt.Println(key)
				if _, ok := m[key]; !ok {

					m[key] = trs
					return
				}
				if len(trs) > 1 {
					var insert []string
				loop:
					for _, tr := range trs[:1] {
						for _, v := range m[key] {
							if v == tr {
								continue loop
							}
						}
						insert = append(insert, tr)
					}
					if len(insert) > 0 {
						m[key] = append(m[key], insert...)
					}

				}

			})
		})
	})
	//num := 0
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("comment3.json", j, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range m {
		fmt.Printf("%s ----> %v\n", i, v)
		// num++
		// if num == 10 {
		// 	break
		// }
	}
}

// 去除重复的
func filter() {
	j, err := ioutil.ReadFile("./comment3.json")
	if err != nil {
		log.Fatal(err)
	}
	var m map[string][]string
	err = json.Unmarshal(j, &m)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range m {
		var str []string
	loop:
		for _, s := range v {
			for _, ss := range str {
				if ss == s {
					continue loop
				}
			}
			str = append(str, s)
		}
		m[k] = str
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("comment4.json", b, 0644)
}

func Text(s *goquery.Selection) string {
	var buf bytes.Buffer

	// Slightly optimized vs calling Each: no single selection object created
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			// Keep newlines and spaces, like jQuery
			buf.WriteString(fmt.Sprintf("%s ",n.Data))
		}
		if n.FirstChild != nil {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	for _, n := range s.Nodes {
		f(n)
	}

	return buf.String()
}
