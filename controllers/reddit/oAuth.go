package reddit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	models "github.com/R3l3ntl3ss/Meme_Api/models/reddit"
)

// GetAccessToken : Get temporary Access Token based on App client ID and Secret
func (r *Reddit) GetAccessToken() (accessToken string) {

	encodedCredentials := r.EncodeCredentials()

	// Reddit URL to get access token
	url := "https://www.reddit.com/api/v1/access_token"

	// Set grant type to client_credentials as POST body
	payload := strings.NewReader("grant_type=client_credentials")

	req, _ := http.NewRequest("POST", url, payload)

	// Set Headers including the User Agent and the Authorization with the encoded credentials
	req.Header.Add("User-Agent", r.UserAgent)
	req.Header.Add("Authorization", "Basic "+encodedCredentials)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "www.reddit.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "29")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	// Make Request
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Error while making connecting to : " + url)
		return ""
	}

	// Close the response body
	defer res.Body.Close()

	// Read the response
	body, _ := ioutil.ReadAll(res.Body)

	var accessTokenBody models.AccessTokenBody

	if err := json.Unmarshal(body, &accessTokenBody); err != nil {
		log.Println("Error while Unmarshalling AccessTokenBody", err)
		return ""
	}

	return accessTokenBody.AccessToken
}
