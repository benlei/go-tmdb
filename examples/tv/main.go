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

	tvShow, err := tmdbClient.GetTVDetails(1399, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tvShow.Name)

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "credits"
	options["language"] = "pt-BR"

	tvShow, err = tmdbClient.GetTVDetails(1399, options)

	if err != nil {
		fmt.Println(err)
	}

	// Credits - Iterate cast from append to response.
	for _, v := range tvShow.TVCreditsAppend.Credits.Cast {
		fmt.Println(v.Name)
	}
}
