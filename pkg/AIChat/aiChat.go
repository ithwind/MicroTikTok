package AIChat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getAccessToken() (string, error) {
	// 使用 API Key，Secret Key 获取access_token，替换下面的应用API Key、应用Secret Key
	clientID := "A9MgIUj2W83dtOKtg6nP1Has"
	clientSecret := "KGlLuUMSlQWu25LoG9axKAHI5085NkRj"

	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret)

	response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("")))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return accessToken, nil
}
func GetResponse(userInput string) (*http.Response, error) {
	accessToken, err := getAccessToken()
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return nil, err
	}

	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant?access_token=" + accessToken

	payloadData := map[string]interface{}{
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": userInput,
			},
		},
	}
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return nil, err
	}

	return response, nil
}
func Chat(userInput string) string {
	response, err := GetResponse(userInput)
	if err != nil {
		fmt.Println("Error getting response:", err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		fmt.Println("Error parsing response JSON:", err)
		return ""
	}

	resultValue, ok := result["result"].(string)
	if !ok {
		fmt.Println("Error extracting 'result' field from response")
		return ""
	}
	return resultValue
}
