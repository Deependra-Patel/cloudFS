package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type File struct {
	alphaNumericCount int
	numWords          int
}

func GetStats(path string) File {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	data := string(dat)
	var file File

	split := strings.Split(data, " ")
	for i := 0; i < len(split); i++ {
		nonAlphanumeric := regexp.MustCompile("[^A-Za-z0-9]")
		alphaNumeric := nonAlphanumeric.ReplaceAllString(split[i], "")
		if alphaNumeric != "" {
			file.numWords ++
		}
		file.alphaNumericCount += len(alphaNumeric)
	}
	return file
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	checkError(r.ParseForm())
	fileName := r.PostForm.Get("file")
	content := r.PostForm.Get("content")
	file, err := os.Create(fileName)
	checkError(err)
	_, err = file.WriteString(content)
	responseCode := http.StatusOK
	if err != nil {
		responseCode = http.StatusInternalServerError
	}
	w.WriteHeader(responseCode)
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	checkError(r.ParseForm())
	fileName := r.PostForm.Get("file")
	bytes, err := ioutil.ReadFile(fileName)
	checkError(err)
	_, err = w.Write(bytes)
	checkError(err)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	checkError(r.ParseForm())
	fileName := r.PostForm.Get("file")
	err := os.Remove(fileName)
	checkError(err)
	responseCode := http.StatusOK
	if err != nil {
		responseCode = http.StatusInternalServerError
	}
	w.WriteHeader(responseCode)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	checkError(r.ParseForm())
	folder := r.PostForm.Get("file")
	_, err := ioutil.ReadDir(folder)
	checkError(err)
}

func main() {
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/replace", createHandler)
	http.HandleFunc("/retrieve", retrieveHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/stats", statsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
