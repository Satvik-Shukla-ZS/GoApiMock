package Parser

import (
	"bufio"
	"strings"
)

type Route struct {
	Path    string
	Method  string
	Entity  string
	Mapping []string
}

func ParseRouteFile(routeFiles []RouteFileContent) (map[string]Route, error) {
	routes := make(map[string]Route)

	for _, routeFile := range routeFiles {
		scanner := bufio.NewScanner(strings.NewReader(routeFile.Content))
		var currentRoute Route

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if currentRoute.Method == "" {
				parts := strings.Split(line, " ")
				if len(parts) >= 2 {
					currentRoute = Route{
						Path:   parts[1],
						Method: parts[0],
					}
				}
			} else if currentRoute.Entity == "" {
				parts := strings.Split(line, " ")
				if len(parts) >= 2 {
					currentRoute.Entity = strings.Trim(parts[0], "")
				}
			} else if len(line) > 0 {
				mapping := strings.TrimSpace(line)
				currentRoute.Mapping = append(currentRoute.Mapping, mapping)
			}

			if len(line) == 0 {
				routes[currentRoute.Path] = currentRoute
				currentRoute = Route{}
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}

		routes[currentRoute.Path] = currentRoute
	}

	return routes, nil
}
