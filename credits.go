package tmdb

import (
	"fmt"
)

// CreditsDetails type is a struct for credits JSON response.
type CreditsDetails struct {
	CreditType string       `json:"credit_type"`
	Department string       `json:"department"`
	Job        string       `json:"job"`
	Media      CreditMedia  `json:"media"`
	MediaType  string       `json:"media_type"`
	ID         string       `json:"id"`
	Person     CreditPerson `json:"person"`
}

type CreditMedia struct {
	Adult            bool                 `json:"adult,omitempty"`          // Movie
	OriginalName     string               `json:"original_name,omitempty"`  // TV
	OriginalTitle    string               `json:"original_title,omitempty"` // Movie
	ID               int64                `json:"id"`
	Name             string               `json:"name,omitempty"` // TV
	VoteCount        int64                `json:"vote_count"`
	VoteAverage      float32              `json:"vote_average"`
	FirstAirDate     string               `json:"first_air_date,omitempty"` // TV
	PosterPath       string               `json:"poster_path"`
	ReleaseDate      string               `json:"release_date,omitempty"` // Movie
	Title            string               `json:"title,omitempty"`        // Movie
	Video            bool                 `json:"video,omitempty"`        // Movie
	GenreIDs         []int64              `json:"genre_ids"`
	OriginalLanguage string               `json:"original_language"`
	BackdropPath     string               `json:"backdrop_path"`
	Overview         string               `json:"overview"`
	OriginCountry    []string             `json:"origin_country,omitempty"` // TV
	Popularity       float32              `json:"popularity"`
	Character        string               `json:"character"`
	Episodes         []CreditMediaEpisode `json:"episodes,omitempty"` // TV
	Seasons          []CreditMediaSeason  `json:"seasons,omitempty"`  // TV
}

type CreditMediaEpisode struct {
	AirDate       string `json:"air_date"`
	EpisodeNumber int64  `json:"episode_number"`
	Name          string `json:"name"`
	Overview      string `json:"overview"`
	SeasonNumber  int    `json:"season_number"`
	StillPath     string `json:"still_path"`
}

type CreditMediaSeason struct {
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
	ShowID       int64  `json:"show_id"`
}

type CreditPerson struct {
	Adult              bool                   `json:"adult"`
	Gender             int                    `json:"gender"`
	Name               string                 `json:"name"`
	ID                 int64                  `json:"id"`
	KnownFor           []CreditPersonKnownFor `json:"known_for"`
	KnownForDepartment string                 `json:"known_for_department"`
	ProfilePath        string                 `json:"profile_path"`
	Popularity         float32                `json:"popularity"`
}

// CreditPersonKnownFor some fields can be both an integer and a float...
// - https://github.com/cyruzin/golang-tmdb/pull/34
// - https://www.themoviedb.org/talk/6026eebd6a3448003e155bbc
// That being said... all referenced credit ids seem to be fixed shrujj
type CreditPersonKnownFor struct {
	Adult            bool     `json:"adult,omitempty"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIDs         []int64  `json:"genre_ids"`
	ID               int64    `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	Title            string   `json:"title,omitempty"`
	Video            bool     `json:"video,omitempty"`
	VoteAverage      float32  `json:"vote_average"`
	VoteCount        int64    `json:"vote_count"`
	Popularity       float32  `json:"popularity"`
	MediaType        string   `json:"media_type"`
	OriginalName     string   `json:"original_name,omitempty"`
	Name             string   `json:"name,omitempty"`
	FirstAirDate     string   `json:"first_air_date,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
}

// GetCreditDetails get a movie or TV credit details by id.
//
// https://developers.themoviedb.org/3/credits/get-credit-details
func (c *Client) GetCreditDetails(
	id string,
) (*CreditsDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%s?api_key=%s",
		baseURL,
		creditURL,
		id,
		c.apiKey,
	)
	creditsDetails := CreditsDetails{}
	if err := c.get(tmdbURL, &creditsDetails); err != nil {
		return nil, err
	}
	return &creditsDetails, nil
}
