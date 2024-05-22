package main

import (
	api "Govrc/api"
	"bufio"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Username Plz :")
	username, _ := reader.ReadString('\n')
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Password Plz :")
	password, _ := reader1.ReadString('\n')

	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}
	api.Login(strings.TrimSpace(username), strings.TrimSpace(password), client)
}
