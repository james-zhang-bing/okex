package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("/home/james/go/src/github.com/okex/models/account/account_models.go")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	st := FindAllStruct(r)
	for _, v := range st {
		for _, line := range v.Line {
			fmt.Println(line)
		}
	}
	sts := toPrefix(st)
	for _, v := range sts {
		fmt.Println(v.structName)
		for _, vv := range v.sf {
			fmt.Printf("%+v\n", vv)
		}
	}
}

var tmpl = `func (m *MaxWithdrawal)String()string{
	var str string
	str=fmt.Sprintf("%s\n[comment]: %v",str,m.Ccy)
	return str
}`

func FindAllStruct(content *bufio.Reader) []*Fields {
	typeStruct := make(map[string][]string)
	var sliceSort []string
	var cur string
	pending := false
	for {
		l, _, err := content.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		line := string(l)
		fields := strings.Fields(line)
		if pending {
			typeStruct[cur] = append(typeStruct[cur], line)
		}
		for i, v := range fields {
			if v == `{` && fields[i-1] == "struct" {
				cur = fields[0]
				sliceSort = append(sliceSort, cur)
				typeStruct[cur] = append(typeStruct[cur], line)
				pending = true
				break
			}
			if v == `}` && pending && i == 0 {
				pending = false
			}
		}
	}
	var Field []*Fields
	for _, v := range sliceSort {
		Field = append(Field, &Fields{
			Name: v,
			Line: typeStruct[v],
		})
	}
	return Field
}

type Fields struct {
	Name string
	Line []string
}

func toPrefix(fs []*Fields) []*structType {
	var sts []*structType
	var sf structField
	for _, v := range fs {
		var st structType
		for _, line := range v.Line[1 : len(v.Line)-1] {
			if strings.Contains(line, "//") {
				sf.comment = getMeans(line)
			} else {
				fields := strings.Fields(line)
				sf.name = fields[0]
			}
			if sf.name != "" {
				st.sf = append(st.sf, sf)
				sf = structField{}
			}
		}
		sts = append(sts, &st)
	}
	return sts
}

type structField struct {
	comment string
	name    string
}

type structType struct {
	structName string
	sf         []structField
}

func getMeans(comment string) string {
	comment = strings.ReplaceAll(comment, "，", " ")
	fields := strings.Fields(comment)
	for _, v := range fields {
		if v == "是" || v == "否" || v == "可选" || v=="//" {
			continue
		}
		if isChinese(v){
			return v
		}
	}
	return "---"
}

func isChinese(s string)bool{
	for _,v:=range s{
		if unicode.Is(unicode.Han,v){
			return true
		}
	}
	return false
}