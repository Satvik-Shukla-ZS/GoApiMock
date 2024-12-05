package Generator

import (
	"GoApiMock/Parser"
	"math/rand"
	"time"
)

type Options struct {
	Min int
	Max int
}

func GenerateRandomString(min, max int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(max-min+1) + min
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRandomEmail(min, max int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(max-min+1) + min
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result) + "@example.com"
}

func GenerateRandomDatetime() string {
	return time.Now().Format("02/01/2006")
}

func GenerateEntities(generator map[string]Options, entities map[string]Parser.Entity) map[string][]map[string]string {
	result := make(map[string][]map[string]string)
	for key, options := range generator {
		if entity, exists := entities[key]; exists {
			length := rand.Intn(options.Max-options.Min+1) + options.Min
			for i := 0; i < length; i++ {
				fieldValues := make(map[string]string)
				for fieldName, field := range entity.Fields {
					switch field.Type {
					case "string":
						fieldValues[fieldName] = GenerateRandomString(field.MinChar, field.MaxChar)
					case "email":
						fieldValues[fieldName] = GenerateRandomEmail(field.MinChar, field.MaxChar)
					case "datetime":
						fieldValues[fieldName] = GenerateRandomDatetime()
						// Add more type handling as needed
					}
				}
				result[key] = append(result[key], fieldValues)
			}
		}
	}
	return result
}
