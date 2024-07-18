package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

const apiURL = "https://api.coinbase.com"
const sandboxURL = "https://api-public.sandbox.exchange.coinbase.com"

var baseURL string
var timeStamp string
var apiKey string
var apiPassPhrase string
var apiSecret string
var apiSecretB64 []byte

func SendRequest(reqMethod, reqPath, reqBody string, queryParams map[string]string, sandbox, verbose bool) (http.Response, []byte, error) {
	configRequestVars(sandbox)

	// Create HTTP Request
	url := baseURL + reqPath
	req, _ := http.NewRequest(reqMethod, url, bytes.NewBuffer([]byte(reqBody)))

	if queryParams != nil {
		q := req.URL.Query()
		for key, value := range queryParams {
			if verbose {
				fmt.Println("Query Param:", key, "=", value)
			}
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	message := timeStamp + reqMethod + req.URL.RequestURI() + reqBody
	signedMessage := signMessage(message)

	// Debug API Credentials
	if verbose {
		fmt.Println("API Base URL:", baseURL)
		fmt.Println("HTTP Request URL", req.URL.String())
		fmt.Println("HTTP Request URI", req.URL.RequestURI())
		fmt.Println("API Key:", apiKey)
		fmt.Println("API Passphrase:", apiPassPhrase)
		fmt.Println("API Secret:", apiSecret)
		fmt.Println("API Secret Decoded:", apiSecretB64)
		fmt.Println("Time Stamp:", timeStamp)
		fmt.Println("Message:", message)
		fmt.Println("Signed Message:", signedMessage)
		fmt.Println("")
	}

	// Set Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("CB-ACCESS-KEY", apiKey)
	req.Header.Add("CB-ACCESS-SIGN", signedMessage)
	req.Header.Add("CB-ACCESS-TIMESTAMP", timeStamp)
	req.Header.Add("CB-ACCESS-PASSPHRASE", apiPassPhrase)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(res, err)
		return *res, nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return *res, body, err
	}
	return *res, body, nil
}

func signMessage(message string) string {
	mac := hmac.New(sha256.New, apiSecretB64)
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func configRequestVars(sandbox bool) {

	// Get Current Time
	now := time.Now()
	timeStamp = strconv.Itoa(int(now.Unix()))
	apiKey = viper.GetString("coinbase.apiKey")
	apiPassPhrase = viper.GetString("coinbase.apiPassPhrase")
	apiSecret = viper.GetString("coinbase.apiSecret")
	apiSecretB64, _ = base64.StdEncoding.DecodeString(apiSecret) // base64 decode API secret

	if sandbox {
		baseURL = sandboxURL
	} else {
		baseURL = apiURL
	}
}
