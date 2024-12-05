package Init

import (
	"GoApiMock/Parser"
)

func InitRouteParsing() (map[string]Parser.Route, error) {
	routesData := Parser.ReadAllRouteFiles()

	parsed, err := Parser.ParseRouteFile(routesData)

	return parsed, err
}
