package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	CheckError(r.ParseForm())
	fileName := r.PostForm.Get("file")
	content := r.PostForm.Get("content")
	file, err := os.Create(fileName)
	CheckError(err)
	_, err = file.WriteString(content)
	responseCode := http.StatusOK
	if err != nil {
		responseCode = http.StatusInternalServerError
	}
	w.WriteHeader(responseCode)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	CheckError(r.ParseForm())
	fileName := r.PostForm.Get("file")
	bytes, err := ioutil.ReadFile(fileName)
	CheckError(err)
	_, err = w.Write(bytes)
	CheckError(err)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	CheckError(r.ParseForm())
	fileName := r.PostForm.Get("file")
	err := os.Remove(fileName)
	CheckError(err)
	responseCode := http.StatusOK
	if err != nil {
		responseCode = http.StatusInternalServerError
	}
	w.WriteHeader(responseCode)
}
