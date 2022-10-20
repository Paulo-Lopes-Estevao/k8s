package main

import "net/http"

func maina() {
	http.HandleFunc("/", HelloWord)
	http.ListenAndServe(":8080", nil)
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Word</h1>"))
}
