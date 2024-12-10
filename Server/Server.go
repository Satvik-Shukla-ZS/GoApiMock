package Server

import (
	"GoApiMock/Parser"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var routes map[string]Parser.Route
var dataContext map[string][]map[string]any

func InitServer(routesPassed map[string]Parser.Route, dataContextPassed map[string][]map[string]any) {
	routes = routesPassed
	dataContext = dataContextPassed

	http.HandleFunc("/", routeHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func extractParams(requestPath, routePath string) map[string]string {
	params := make(map[string]string)
	requestParts := strings.Split(requestPath, "/")
	routeParts := strings.Split(routePath, "/")

	for i, part := range routeParts {
		if strings.HasPrefix(part, "<>") && i < len(requestParts) {
			paramName := strings.Trim(part, "<>")
			params[paramName] = requestParts[i]
		}
	}

	return params
}

func isValidURL(routes map[string]Parser.Route, url string) (Parser.Route, bool) {
	for routePath, route := range routes {
		if matchURL(routePath, url) {
			return route, true
		}
	}
	return Parser.Route{}, false
}

func matchURL(routePath, url string) bool {
	routeParts := strings.Split(routePath, "/")
	urlParts := strings.Split(url, "/")

	if len(routeParts) != len(urlParts) {
		return false
	}

	for i, part := range routeParts {
		if strings.HasPrefix(part, "<>") {
			continue
		}
		if part != urlParts[i] {
			return false
		}
	}

	return true
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	route, exists := isValidURL(routes, r.URL.Path)

	if !exists {
		http.NotFound(w, r)
		return
	}

	entityName := route.Entity
	isMultiple := false
	var limit int

	// Check for suffixes [] or [3]
	re := regexp.MustCompile(`\[(\d*)\]$`)
	matches := re.FindStringSubmatch(entityName)
	if len(matches) > 0 {
		if matches[1] == "" {
			isMultiple = true
		} else {
			limit, _ = strconv.Atoi(matches[1])
		}
		entityName = entityName[:len(entityName)-len(matches[0])]
	}

	entityData, exists := dataContext[entityName]

	if !exists {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}

	filteredData := []map[string]any{}
	for _, data := range entityData {
		tempData := make(map[string]any)
		for _, keys := range route.Mapping {
			tempData[keys] = data[keys]
		}
		filteredData = append(filteredData, tempData)
		if !isMultiple && limit == 0 {
			break
		}
		if limit > 0 && len(filteredData) >= limit {
			break
		}
	}

	response := map[string]interface{}{}
	if isMultiple || limit > 0 {
		response["data"] = filteredData
	} else {
		response["data"] = filteredData[0]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
