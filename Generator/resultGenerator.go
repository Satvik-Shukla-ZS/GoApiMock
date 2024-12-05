package Generator

import (
	"GoApiMock/Parser"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

type Options struct {
	Min  int
	Max  int
	Name string
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

func GenerateRandomContent(min, max int) string {
	const charset = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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

func GenerateRandomUUID() string {
	return uuid.New().String()
}

func GenerateRandomEnum(options []string) string {
	rand.Seed(time.Now().UnixNano())
	return options[rand.Intn(len(options))]
}

func GenerateRandomBool() bool {
	return rand.Intn(2) == 1
}

func GenerateRandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

func GenerateRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func GenerateEntities(generator []Options, entities map[string]Parser.Entity) map[string][]map[string]any {
	result := make(map[string][]map[string]any)
	for _, options := range generator {
		if entity, exists := entities[options.Name]; exists {
			length := rand.Intn(options.Max-options.Min+1) + options.Min
			for i := 0; i < length; i++ {
				fieldValues := make(map[string]any)
				for fieldName, field := range entity.Fields {
					switch field.Type {
					case "string":
						fieldValues[fieldName] = GenerateRandomString(field.MinChar, field.MaxChar)
					case "email":
						fieldValues[fieldName] = GenerateRandomEmail(field.MinChar, field.MaxChar)
					case "content":
						fieldValues[fieldName] = GenerateRandomContent(field.MinChar, field.MaxChar)
					case "datetime":
						fieldValues[fieldName] = GenerateRandomDatetime()
					case "UUID":
						fieldValues[fieldName] = GenerateRandomUUID()
					case "enum":
						fieldValues[fieldName] = GenerateRandomEnum(field.EnumOptions)
					case "bool":
						fieldValues[fieldName] = GenerateRandomBool()
					case "float":
						fieldValues[fieldName] = GenerateRandomFloat(float64(field.MinChar), float64(field.MaxChar))
					case "int":
						fieldValues[fieldName] = GenerateRandomInt(field.MinChar, field.MaxChar)
					default:
						// Handle dynamic types like USER.id
						if strings.Contains(field.Type, ".") {
							parts := strings.Split(field.Type, ".")
							if len(parts) == 2 {
								entityType, entityField := parts[0], parts[1]
								if relatedEntities, found := result[entityType]; found && len(relatedEntities) > 0 {
									randomIndex := rand.Intn(len(relatedEntities))
									fieldValues[fieldName] = relatedEntities[randomIndex][entityField]
								}
							}
						}
					}
				}
				result[options.Name] = append(result[options.Name], fieldValues)
			}
		}
	}
	return result
}
