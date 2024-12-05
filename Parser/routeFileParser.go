package Parser

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

type RouteFileContent struct {
	Name    string
	Content string
}

func ReadAllRouteFiles() []RouteFileContent {
	files, err := filepath.Glob("./Schema/*.route.sch")
	if err != nil {
		log.Fatal(err)
	}

	routeFiles := make([]RouteFileContent, 0)
	re := regexp.MustCompile(`^\d+\.route\.sch$`)

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
