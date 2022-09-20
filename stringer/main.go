package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"unicode"
)

func main() {

	ps := scanFile("/home/james/go/src/github.com/okex/models")
	for _, p := range ps {
		Stringer(p)
	}

}

var tmpl = `func (m *{{ .StructName }})String()string{
	var str string
	{{ range .Sf }}str=fmt.Sprintf("%s\n{{ .Comment }}: %v",str,m.{{ .Name }})
	{{ end }}return str
}`
var noCommentTmpl = `func (m *{{ .StructName }})String()string{
	var str string
	{{ range .Sf }}str=fmt.Sprintf("%s\n{{ .Name }}:%v",str,m.{{ .Name }})
	{{ end }}return str
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
		st.StructName = v.Name
		for _, line := range v.Line[1 : len(v.Line)-1] {
			if strings.Contains(line, "//") {
				sf.Comment = getMeans(line)
			} else {
				fields := strings.Fields(line)
				sf.Name = fields[0]
			}
			if sf.Name != "" {
				st.Sf = append(st.Sf, sf)
				sf = structField{}
			}
		}
		sts = append(sts, &st)
	}
	return sts
}

func splitNoComment(sts []*structType) (comment, noComment []*structType) {
loop:
	for _, v := range sts {
		for _, vv := range v.Sf {
			if vv.Comment != "" {
				comment = append(comment, v)
				continue loop
			}
		}
		noComment = append(noComment, v)
	}
	return 
}

type structField struct {
	Comment string
	Name    string
}

type structType struct {
	StructName string
	Sf         []structField
}

func getMeans(comment string) string {
	comment = strings.ReplaceAll(comment, "，", " ")
	fields := strings.Fields(comment)
	for _, v := range fields {
		if v == "是" || v == "否" || v == "可选" || v == "//" {
			continue
		}
		if isChinese(v) {
			return v
		}
	}
	return "---"
}

func isChinese(s string) bool {
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}

func Stringer(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	st := FindAllStruct(r)

	stss := toPrefix(st)
	sts,noComment:=splitNoComment(stss)
	f.Close()
	f, err = os.Open(filePath)
	fatalErr(err)
	read := bufio.NewReader(f)
	var lines []string
	for {
		line, _, err := read.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}
	f.Close()

	err = os.Remove(filePath)
	fatalErr(err)

	f, err = os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer f.Sync()
	w := bufio.NewWriter(f)
	defer w.Flush()
	_, err = w.WriteString(lines[0])
	fatalErr(err)
	fmt.Printf("import \"fmt\"\n")
	_, err = w.WriteString(fmt.Sprintf("\nimport \"fmt\""))
	fatalErr(err)
	for _, line := range lines[1:] {
		fmt.Printf("\n%s", line)
		_, err := w.WriteString(fmt.Sprintf("\n%s", line))
		fatalErr(err)
	}
	t := template.New("stringer")
	t, err = t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteString("\n")
	for _, v := range sts {
		err := t.Execute(w, *v)
		fatalErr(err)
		_, err = w.WriteString("\n")
		fatalErr(err)
	}
	simple:=template.New("nocomment")
	simple,err=simple.Parse(noCommentTmpl)
	fatalErr(err)
	for _, v := range noComment {
		err := simple.Execute(w, *v)
		fatalErr(err)
		_, err = w.WriteString("\n")
		fatalErr(err)
	}
}

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
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
