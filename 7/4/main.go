package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] + " at " + substrs[1])
}

func main() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname exapmle")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`
    hello {{.UserName}}!
    {{range .Emails}}
      an email {{.|emailDeal}}
    {{end}}
    {{with .Friends}}
    {{range .}}
      my friend name is {{.Fanme}}
    {{end}}
    {{end}}
  `)
	p := Person{
		UserName: "Astaxie",
		Emails:   []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)

	fmt.Println("\n=== 条件分岐 ===")

	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空のpipeline: {{if ``}}出力されない{{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("空でないpipeline: {{if `anything`}}出力される{{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else: {{if `anything`}}出力される{{else}}出力されない{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}
