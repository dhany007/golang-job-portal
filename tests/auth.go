package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func GetAccessToken(body string) (result string, err error) {
	db, err := InitPostgresTest()
	if err != nil {
		return "", err
	}

	server := NewServer(db)
	response := server.Request(http.MethodPost, "http://localhost:11010/users/login", strings.NewReader(body))

	temp, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(temp, &responseBody)

	result = responseBody["data"].(map[string]interface{})["access_token"].(string)

	return result, nil
}
