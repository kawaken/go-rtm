package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

// Param is param
type Param struct {
	GoName  string
	APIName string
}

// Method is method
type Method struct {
	Name     string
	FullName string
	Requires []*Param
}

// Arguments returns method arguments.
func (m *Method) Arguments() string {
	nt := []string{}
	for _, p := range m.Requires {
		nt = append(nt, fmt.Sprintf("%s string", p.GoName))
	}

	return strings.Join(nt, ", ")
}

// Package is package
type Package struct {
	Name    string
	Methods []*Method
}

func writeFile(tmpl *template.Template, fullName string, params string) error {

	elms := strings.Split(fullName, ".")
	if len(elms) < 3 {
		return fmt.Errorf("fullName is invalid %s", fullName)
	}
	p := &Package{Name: elms[1]}

	name := snaker.SnakeToCamel(strings.Join(elms[2:], "_"))
	m := &Method{FullName: fullName, Name: name}

	elms = strings.Split(params, ",")
	var r []*Param
	for _, e := range elms {
		if e == "" {
			continue
		}

		gn := []byte(snaker.SnakeToCamel(e))
		gn[0] = gn[0] | ('a' - 'A')

		r = append(r, &Param{
			APIName: e,
			GoName:  string(gn),
		})
	}
	m.Requires = r

	p.Methods = []*Method{m}

	err := os.MkdirAll(p.Name, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := snaker.CamelToSnake(m.Name)
	f, err := os.Create(fmt.Sprintf("%s/%s.go", p.Name, fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, p)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func main() {
	tmpl := template.Must(template.ParseFiles("./file.tmpl"))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}

		params := ""
		if len(fields) > 1 {
			params = fields[1]
		}

		err := writeFile(tmpl, fields[0], params)
		if err != nil {
			fmt.Printf("write error: %s\n", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("read error: %s\n", err)
		return
	}

}
