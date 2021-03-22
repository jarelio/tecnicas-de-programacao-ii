package utils

import "fmt"

func ResultMessageAndDataToJSON(message, data string) string {
	return fmt.Sprintf("{\"result\": {\"message\": \"%s\", \"data\": %s}}", message, data)
}

func ErrorMessageToJSON(message string) string {
	return fmt.Sprintf("{\"error\": \"%s\"}", message)
}
