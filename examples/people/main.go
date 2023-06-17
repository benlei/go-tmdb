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

	person, err := tmdbClient.GetPersonDetails(117642, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person.Name)

	fmt.Println("------")

	//With options
	options := make(map[string]string)
	options["append_to_response"] = "images"
	options["language"] = "pt-BR"

	person, err = tmdbClient.GetPersonDetails(117642, options)

	if err != nil {
		fmt.Println(err)
	}

	// Images - Iterate profiles from append to response.
	for _, v := range person.PersonImagesAppend.Images.Profiles {
		fmt.Println(v.FilePath)
	}
}
