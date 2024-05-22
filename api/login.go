package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

func login(username string, password string) {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}
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
	fmt.Printf("%s\n", body)
	if strings.Contains(string(body), "mail") {
		emailapiurl := "https://api.vrchat.cloud/api/1/auth/twofactorauth/emailotp/verify"
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Email code Plz : \n")
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
	} else if strings.Contains(string(body), "app") {
		fmt.Printf("App code Plz : \n")
	}
	return
}

func main() {
	login("violet404", "Amber7300.")
}
