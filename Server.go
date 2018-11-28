package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileStat struct {
	alphaNumericWordLengths []int
}

func GetStats(path string) FileStat {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	data := string(dat)
	var stat FileStat

	split := strings.Split(data, " ")
	for i := 0; i < len(split); i++ {
		nonAlphanumeric := regexp.MustCompile("[^A-Za-z0-9]")
		alphaNumeric := nonAlphanumeric.ReplaceAllString(split[i], "")
		if alphaNumeric != "" {
			stat.alphaNumericWordLengths = append(stat.alphaNumericWordLengths, len(alphaNumeric))
		}
	}
	return stat
}

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

func statsHandler(w http.ResponseWriter, r *http.Request) {
	CheckError(r.ParseForm())
	folder := r.PostForm.Get("file")
	var stats []FileStat
	var totalNumBytes int64
	err := filepath.Walk(folder, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() { // Skip the directories
			return nil
		}

		fmt.Printf("Visited: %s\n", path)
		stats = append(stats, GetStats(path))
		totalNumBytes += f.Size()
		return nil
	})
	CheckError(err)
	_, err = w.Write([]byte("Total number of files: " + fmt.Sprintf("%d", len(stats)) + "\n"))
	CheckError(err)

	var alphanumericCharsPerFile []int
	var allWordLengths []int
	for index := range stats {
		wordLengthsInFile := stats[index].alphaNumericWordLengths
		alphanumericCharsPerFile = append(alphanumericCharsPerFile, Sum(wordLengthsInFile))
		allWordLengths = append(allWordLengths, wordLengthsInFile...)
	}

	average := Avg(alphanumericCharsPerFile)
	_, err = w.Write([]byte("Average number of alphanumeric characters per text file: " +
		FloatToString(average) + ", StdDev: " + FloatToString(StdDev(alphanumericCharsPerFile, average)) + "\n"))
	CheckError(err)

	average = Avg(allWordLengths)
	_, err = w.Write([]byte("Average word length in folder: " +
		FloatToString(average) + ", StdDev: " + FloatToString(StdDev(allWordLengths, average)) + "\n"))
	CheckError(err)

	_, err = w.Write([]byte("Total number of bytes stored in folder: " + IntToString(totalNumBytes) + "\n"))
	CheckError(err)
}

func main() {
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/replace", createHandler)
	http.HandleFunc("/retrieve", retrieveHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/stats", statsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
