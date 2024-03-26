package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")

	data := map[string]string{
		"name": "codewf",
		"someStr": "这是一个开始",
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", myWeb)

	fmt.Println("Server will start, please browser http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server start error: ", err)
	}
}