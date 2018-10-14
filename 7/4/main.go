package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string
}

func main() {
	t := template.New("fieldname exapmle")
	t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
	p := Person{UserName: "Astaxie", email: "test@test.com"}
	t.Execute(os.Stdout, p)
}
