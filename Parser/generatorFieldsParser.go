package Parser

import (
	"fmt"
	"strconv"
	"strings"
)

type GeneratorOptions struct {
	Min  int
	Max  int
	Name string
}

func ParseGeneratorData(data []RouteFileContent) ([]GeneratorOptions, error) {
	var options []GeneratorOptions

	for _, fileContent := range data {
		lines := strings.Split(fileContent.Content, "\n")

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Split(line, ":")
			if len(parts) < 3 {
				return nil, fmt.Errorf("invalid format: %s", line)
			}

			name := strings.TrimSpace(parts[0])
			minStr := strings.TrimSpace(strings.Split(parts[1], "=")[1])
			maxStr := strings.TrimSpace(strings.Split(parts[2], "=")[1])

			minVal, err := strconv.Atoi(minStr)
			if err != nil {
				return nil, fmt.Errorf("invalid min value: %s", minStr)
			}

			maxVal, err := strconv.Atoi(maxStr)
			if err != nil {
				return nil, fmt.Errorf("invalid max value: %s", maxStr)
			}

			options = append(options, GeneratorOptions{
				Min:  minVal,
				Max:  maxVal,
				Name: name,
			})
		}
	}

	return options, nil
}
