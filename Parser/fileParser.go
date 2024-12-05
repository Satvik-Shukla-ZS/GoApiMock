package Parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type FileContent struct {
	Name    string
	Content string
}

func ReadAllEntityFile() []FileContent {
	files, err := filepath.Glob("./Schema/*.entity.sch")
	if err != nil {
		fmt.Println(err)
	}

	fileContents := make([]FileContent, 0)

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		fileObj := FileContent{Name: file, Content: string(data)}
		fileContents = append(fileContents, fileObj)
	}

	return fileContents
}
