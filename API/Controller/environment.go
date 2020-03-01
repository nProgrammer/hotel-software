package Controller

import (
	"fmt"
	"os"
)

func ConfigureEnvironmentVar() (string, string, string, string, string) {
	userLogin := os.Getenv("USER_LOGIN")
	fmt.Println(userLogin)
	userPassword := os.Getenv("USER_PASSWORD")
	fmt.Println(userPassword)
	dbUrl := os.Getenv("DB_URL")
	fmt.Println(dbUrl)
	passM := os.Getenv("MANAGER_PASSWORD")
	fmt.Println(passM)
	token := os.Getenv("URL_TOKEN")
	fmt.Println(token)

	return userLogin, userPassword, dbUrl, passM, token
}
