package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"

	tmdb "github.com/benlei/go-tmdb"
)

func main() {
	godotenv.Load()
	tmdbClient, err := tmdb.Init(os.Getenv("API_KEY"))

	// Setting a custom config for http.Client.
	tmdbClient.SetClientConfig(http.Client{Timeout: time.Second * 5})

	if err != nil {
		fmt.Println(err)
	}

	movie, err := tmdbClient.GetMovieDetails(299536, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie.Title)

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "credits"
	options["language"] = "pt-BR"

	movie, err = tmdbClient.GetMovieDetails(299536, options)

	if err != nil {
		fmt.Println(err)
	}

	// Credits - Iterate cast from append to response.
	for _, v := range movie.MovieCreditsAppend.Credits.Cast {
		fmt.Println(v.Name)
	}
}
