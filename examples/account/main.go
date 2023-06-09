package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"

	tmdb "github.com/benlei/go-tmdb/v2"
)

func main() {
	godotenv.Load()
	tmdbClient, err := tmdb.Init(os.Getenv("API_KEY"))

	if err != nil {
		fmt.Println(err)
	}

	if err := tmdbClient.SetSessionID(os.Getenv("SESSION_ID")); err != nil {
		fmt.Println(err)
	}

	account, err := tmdbClient.GetAccountDetails()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.ID)
	fmt.Println(account.Username)
	fmt.Println(account.Name)
}
