package tmdb

import "fmt"

// GenreList type is a struct for genres movie list JSON response.
type GenreList struct {
	Genres []*Genre `json:"genres"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetGenreMovieList get the list of official genres for movies.
//
// https://developers.themoviedb.org/3/genres/get-movie-list
func (c *Client) GetGenreMovieList(
	urlOptions map[string]string,
) (*GenreList, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie/list?api_key=%s%s",
		baseURL,
		genreURL,
		c.apiKey,
		options,
	)
	genreMovieList := GenreList{}
	if err := c.get(tmdbURL, &genreMovieList); err != nil {
		return nil, err
	}
	return &genreMovieList, nil
}

// GetGenreTVList get the list of official genres for TV shows.
//
// https://developers.themoviedb.org/3/genres/get-tv-list
func (c *Client) GetGenreTVList(
	urlOptions map[string]string,
) (*GenreList, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv/list?api_key=%s%s",
		baseURL,
		genreURL,
		c.apiKey,
		options,
	)
	genreTVList := GenreList{}
	if err := c.get(tmdbURL, &genreTVList); err != nil {
		return nil, err
	}
	return &genreTVList, nil
}
