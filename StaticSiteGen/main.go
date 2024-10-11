package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("public"))
	if err := http.ListenAndServe(":8888", fs); err != nil {
		panic("failed starting server " + err.Error())
	}
}
