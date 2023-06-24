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

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "watch/providers"

	movie, err := tmdbClient.GetMovieDetails(299536, options)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range movie.MovieWatchProvidersAppend.WatchProviders.Results {
		fmt.Println(v.FlatRate)
	}
}
