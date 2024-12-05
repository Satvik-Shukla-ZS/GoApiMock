package Parser

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

type GeneratorFileContent struct {
	Name    string
	Content string
}

func ReadAllGeneratorFiles() []RouteFileContent {
	files, err := filepath.Glob("./Schema/*.generator.sch")
	if err != nil {
		log.Fatal(err)
	}

	routeFiles := make([]RouteFileContent, 0)
	re := regexp.MustCompile(`^\d+\.generator\.sch$`)

	for _, file := range files {
		if re.MatchString(filepath.Base(file)) {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			routeFile := RouteFileContent{Name: file, Content: string(data)}
			routeFiles = append(routeFiles, routeFile)
		}
	}

	return routeFiles
}
