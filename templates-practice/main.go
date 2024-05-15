package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	tmpl := "Hello {{.Name}}!"
	user := User{Name: "James"}
	t, err := template.New("test").Parse(tmpl)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}

}
