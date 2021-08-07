package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name       string
		Age        int
		Occupation string
		City       string
		Interests  [3]string
		Online     bool
	}{"Daniel DeWeerd", 34, "Software Developer", "Seattle", [3]string{"coding", "books", "games"}, true}

	t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
