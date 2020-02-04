package Controller

import (
	"fmt"
	"os"
)

func ConfigureEnvironmentVar() (string, string, string, string, string, string, string) {
	userLogin := os.Getenv("USER_LOGIN")
	fmt.Println(userLogin)
	userPassword := os.Getenv("USER_PASSWORD")
	fmt.Println(userPassword)
	url := os.Getenv("HOST_URL")
	fmt.Println(url)
	dbUrl := os.Getenv("DB_URL")
	fmt.Println(dbUrl)
	passM := os.Getenv("MANAGER_PASSWORD")
	fmt.Println(passM)
	token := os.Getenv("URL_TOKEN")
	fmt.Println(token)
	aUrl := os.Getenv("API_URL")
	fmt.Println(aUrl)

	return userLogin, userPassword, url, dbUrl, passM, token, aUrl
}
