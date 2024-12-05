package main

import (
	"GoApiMock/Generator"
	"GoApiMock/Parser"
	"fmt"
)

func main() {
	data := Parser.ReadAllEntityFile()
	dataParsed := Parser.ParseFileContents(data)

	requested := make(map[string]Generator.Options)
	requested["USER"] = Generator.Options{Min: 5, Max: 10}

	result := Generator.GenerateEntities(requested, dataParsed)
	fmt.Println(result)
}
