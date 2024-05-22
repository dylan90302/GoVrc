package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Login(username string, password string, client http.Client) {
	baseurl := "https://api.vrchat.cloud/api/1/auth/user"
	req, _ := http.NewRequest("GET", baseurl, nil)
	req.SetBasicAuth(username, password)
	req.Header.Set("User-Agent", "VrcGo/0.1.0 jdylan70@gmail.com")
	res, err := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 200 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	type respjsn struct {
		Resp []string `json:"requiresTwoFactorAuth"`
	}
	jsn := respjsn{}
	err = json.Unmarshal(body, &jsn)
	if err != nil {
		fmt.Println(err)
	}
	if jsn.Resp == nil {
		return
	}
	if jsn.Resp[0] == "emailOtp" {
		emailapiurl := "https://api.vrchat.cloud/api/1/auth/twofactorauth/emailotp/verify"
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Email code Plz :")
		code, _ := reader.ReadString('\n')
		fmt.Println(code)
		var codejsn = []byte(`{"code":"` + strings.TrimSpace(code) + `"}`)
		fmt.Println(string(codejsn))
		reqa, _ := http.NewRequest("POST", emailapiurl, bytes.NewBuffer(codejsn))
		reqa.Header.Set("Content-Type", "application/json")
		reqa.Header.Set("User-Agent", "VrcGo/0.1.0 jdylan70@gmail.com")
		resp, _ := client.Do(reqa)
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Printf("Respones : %s", body)
	} else if jsn.Resp[0] == "totp" {
		fa2apiurl := "https://api.vrchat.cloud/api/1/auth/twofactorauth/totp/verify"
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("App code Plz : \n")
		code, _ := reader.ReadString('\n')
		fmt.Println(code)
		var codejsn = []byte(`{"code":"` + strings.TrimSpace(code) + `"}`)
		fmt.Println(string(codejsn))
		reqa, _ := http.NewRequest("POST", fa2apiurl, bytes.NewBuffer(codejsn))
		reqa.Header.Set("Content-Type", "application/json")
		reqa.Header.Set("User-Agent", "VrcGo/0.1.0 jdylan70@gmail.com")
		resp, _ := client.Do(reqa)
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Printf("Respones : %s", body)
	}
	Login(username, password, client)
}
