package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"

	tmdb "github.com/benlei/go-tmdb"
)

func main() {
	godotenv.Load()
	tmdbClient, err := tmdb.Init(os.Getenv("API_KEY"))

	if err != nil {
		fmt.Println(err)
	}

	movie, err := tmdbClient.GetMovieDetails(155, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie.Title)

	fmt.Println("------")

	//With options
	options := make(map[string]string)
	options["append_to_response"] = "videos"
	//options["language"] = "pt-BR"

	movie, err = tmdbClient.GetMovieDetails(155, options)

	if err != nil {
		fmt.Println(err)
	}

	// Images - Iterate profiles from append to response.
	fmt.Printf("%v\n", movie.MovieVideosAppend.Videos.MovieVideosResults)
}
