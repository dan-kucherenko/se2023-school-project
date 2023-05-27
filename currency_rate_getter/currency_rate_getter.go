package currency_rate_getter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// structure for response decoding
type JSONResponse struct {
	Time string  `json:"time"`
	Rate float64 `json:"rate"`
}

// function to get the rate
func GetRateBtcToUah() (string, float64, error) {
	client := http.Client{}
	apiKey := os.Getenv("API_KEY")
	url := "https://rest.coinapi.io/v1/exchangerate/BTC/UAH"

	// create a request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-CoinAPI-Key", apiKey)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", -1, err
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error sending request:", err)
		return "", -1, err
	}
	defer resp.Body.Close()

	// read the response body
	var result JSONResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return "", -1, err
	}
	return result.Time, result.Rate, nil
}
