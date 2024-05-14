package main

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

func encodelogin(username string, password string) string {
	auth := username + ":" + password
	return b64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+encodelogin("username1", "password123"))
	return nil
}

func main() {
	baseurl := "https://api.vrchat.cloud/api/1/auth/user"
	req, _ := http.NewRequest("GET", baseurl, nil)
	req.SetBasicAuth("barrtya", "cyrus90302")
	req.Header.Set("User-Agent", "VrcGo/0.1.0 jdylan70@gmail.com")
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 200 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
