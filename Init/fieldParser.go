package Init

import (
	"GoApiMock/Generator"
	"GoApiMock/Parser"
	"fmt"
)

func InitFieldParsing() map[string][]map[string]any {
	data := Parser.ReadAllEntityFile()

	dataParsed, _ := Parser.ParseFileContents(data)

	generatorFiles := Parser.ReadAllGeneratorFiles()
	formatedGenerator, err := Parser.ParseGeneratorData(generatorFiles)

	if err != nil {
		fmt.Println("Error parsing generator data", err)
		return nil
	}

	requested := make([]Generator.Options, 0)

	for _, entity := range formatedGenerator {
		requested = append(requested, Generator.Options{
			Min:  entity.Min,
			Max:  entity.Max,
			Name: entity.Name,
		})
	}

	result := Generator.GenerateEntities(requested, dataParsed)

	//for key, entity := range result {
	//	fmt.Println(key)
	//	for _, item := range entity {
	//		for keyEntity, value := range item {
	//			fmt.Println("    ", keyEntity, value)
	//		}
	//		fmt.Println()
	//	}
	//}

	return result
}
