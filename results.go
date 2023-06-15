package tmdb

// AccountCreatedListsResults Result Types
type AccountCreatedListsResults struct {
	Results []AccountCreatedListsResult `json:"results"`
}

type AccountCreatedListsResult struct {
	Description   string `json:"description"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	ItemCount     int64  `json:"item_count"`
	LanguageCode  string `json:"iso_639_1"`
	ListType      string `json:"list_type"`
	Name          string `json:"name"`
	PosterPath    string `json:"poster_path"`
}

// AccountFavoriteMoviesResults Result Types
type AccountFavoriteMoviesResults struct {
	Results []AccountFavoriteMoviesResult `json:"results"`
}

type AccountFavoriteMoviesResult struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int64   `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	PosterPath       string  `json:"poster_path"`
	Popularity       float64 `json:"popularity"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int64   `json:"vote_count"`
}

// AccountFavoriteTVShowsResults Result Types
type AccountFavoriteTVShowsResults struct {
	Results []AccountFavoriteTVShowsResult `json:"results"`
}

type AccountFavoriteTVShowsResult struct {
	BackdropPath     string   `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int64  `json:"genre_ids"`
	ID               int64    `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	OriginCountry    []string `json:"origin_country"`
	PosterPath       string   `json:"poster_path"`
	Popularity       float64  `json:"popularity"`
	Name             string   `json:"name"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int64    `json:"vote_count"`
}

// AccountRatedTVEpisodesResults Result Types
type AccountRatedTVEpisodesResults struct {
	Results []AccountRatedTVEpisodesResult `json:"results"`
}

type AccountRatedTVEpisodesResult struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int64   `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int64   `json:"vote_count"`
	Rating         float32 `json:"rating"`
}

// ChangesMovieResults Result Types
type ChangesMovieResults struct {
	Results []ChangesMovieResult `json:"results"`
}

type ChangesMovieResult struct {
	ID    int64 `json:"id"`
	Adult bool  `json:"adult"`
}

// CompanyAlternativeNamesResults Result Types
type CompanyAlternativeNamesResults struct {
	Results []CompanyAlternativeNamesResult `json:"results"`
}

type CompanyAlternativeNamesResult struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// DiscoverMovieResults Result Types
type DiscoverMovieResults struct {
	Results []DiscoverMovieResult `json:"results"`
}

type DiscoverMovieResult struct {
	VoteCount        int64   `json:"vote_count"`
	ID               int64   `json:"id"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	Title            string  `json:"title"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	GenreIDs         []int64 `json:"genre_ids"`
	BackdropPath     string  `json:"backdrop_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
}

// DiscoverTVResults Result Types
type DiscoverTVResults struct {
	Results []DiscoverTVResult `json:"results"`
}

type DiscoverTVResult struct {
	OriginalName     string   `json:"original_name"`
	GenreIDs         []int64  `json:"genre_ids"`
	Name             string   `json:"name"`
	Popularity       float32  `json:"popularity"`
	OriginCountry    []string `json:"origin_country"`
	VoteCount        int64    `json:"vote_count"`
	FirstAirDate     string   `json:"first_air_date"`
	BackdropPath     string   `json:"backdrop_path"`
	OriginalLanguage string   `json:"original_language"`
	ID               int64    `json:"id"`
	VoteAverage      float32  `json:"vote_average"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
}

// SearchCompaniesResults Result Types
type SearchCompaniesResults struct {
	Results []SearchCompaniesResult `json:"results"`
}

type SearchCompaniesResult struct {
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// SearchCollectionsResults Result Types
type SearchCollectionsResults struct {
	Results []SearchCollectionsResult `json:"results"`
}

type SearchCollectionsResult struct {
	Adult            bool   `json:"adult"`
	BackdropPath     string `json:"backdrop_path"`
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	OriginalLanguage string `json:"original_language"`
	OriginalName     string `json:"original_name"`
	Overview         string `json:"overview"`
	PosterPath       string `json:"poster_path"`
}

// SearchKeywordsResults Result Types
type SearchKeywordsResults struct {
	Results []SearchKeywordsResult `json:"results"`
}

type SearchKeywordsResult struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// SearchMoviesResults Result Types
type SearchMoviesResults struct {
	Results []SearchMoviesResult `json:"results"`
}

type SearchMoviesResult struct {
	VoteCount        int64   `json:"vote_count"`
	ID               int64   `json:"id"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	Title            string  `json:"title"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	GenreIDs         []int64 `json:"genre_ids"`
	BackdropPath     string  `json:"backdrop_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
}

// SearchMultiResults Result Types
type SearchMultiResults struct {
	Results []SearchMultiResult `json:"results"`
}

type SearchMultiResult struct {
	PosterPath       string                `json:"poster_path,omitempty"`
	Popularity       float32               `json:"popularity"`
	ID               int64                 `json:"id"`
	Overview         string                `json:"overview,omitempty"`
	BackdropPath     string                `json:"backdrop_path,omitempty"`
	VoteAverage      float32               `json:"vote_average,omitempty"`
	MediaType        string                `json:"media_type"`
	FirstAirDate     string                `json:"first_air_date,omitempty"`
	OriginCountry    []string              `json:"origin_country,omitempty"`
	GenreIDs         []int64               `json:"genre_ids,omitempty"`
	OriginalLanguage string                `json:"original_language,omitempty"`
	VoteCount        int64                 `json:"vote_count,omitempty"`
	Name             string                `json:"name,omitempty"`
	OriginalName     string                `json:"original_name,omitempty"`
	Adult            bool                  `json:"adult,omitempty"`
	ReleaseDate      string                `json:"release_date,omitempty"`
	OriginalTitle    string                `json:"original_title,omitempty"`
	Title            string                `json:"title,omitempty"`
	Video            bool                  `json:"video,omitempty"`
	ProfilePath      string                `json:"profile_path,omitempty"`
	KnownFor         []SearchMultiKnownFor `json:"known_for,omitempty"`
}

type SearchMultiKnownFor struct {
	ID               int64   `json:"id"`
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int64 `json:"genre_ids"`
	MediaType        string  `json:"media_type"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	Popularity       float32 `json:"popularity"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int64   `json:"vote_count"`
}

// SearchPeopleResults Result Types
type SearchPeopleResults struct {
	Results []SearchPeopleResult `json:"results"`
}

type SearchPeopleResult struct {
	Adult              bool                   `json:"adult"`
	Gender             int                    `json:"gender,omitempty"`
	ID                 int64                  `json:"id"`
	KnownFor           []SearchPeopleKnownFor `json:"known_for"`
	KnownForDepartment string                 `json:"known_for_department"`
	Name               string                 `json:"name"`
	Popularity         float32                `json:"popularity"`
	ProfilePath        string                 `json:"profile_path"`
}

type SearchPeopleKnownFor struct {
	Adult            bool     `json:"adult,omitempty"` // Movie
	BackdropPath     string   `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date,omitempty"` // TV
	GenreIDs         []int64  `json:"genre_ids"`
	ID               int64    `json:"id"`
	MediaType        string   `json:"media_type"`
	Name             string   `json:"name,omitempty"` // TV
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name,omitempty"`  // TV
	OriginalTitle    string   `json:"original_title,omitempty"` // Movie
	OriginCountry    []string `json:"origin_country,omitempty"` // TV
	Overview         string   `json:"overview"`
	Popularity       float32  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date,omitempty"` // Movie
	Title            string   `json:"title,omitempty"`        // Movie
	Video            bool     `json:"video,omitempty"`        // Movie
	VoteAverage      float32  `json:"vote_average"`
	VoteCount        int64    `json:"vote_count"`
}

// SearchTVShowsResults Result Types
type SearchTVShowsResults struct {
	Results []SearchTVShowsResult `json:"results"`
}

type SearchTVShowsResult struct {
	ID               int64    `json:"id"`
	BackdropPath     string   `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIDs         []int64  `json:"genre_ids"`
	Name             string   `json:"name"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	OriginCountry    []string `json:"origin_country"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	Popularity       float32  `json:"popularity"`
	VoteAverage      float32  `json:"vote_average"`
	VoteCount        int64    `json:"vote_count"`
}

// TrendingResults Result Types
type TrendingResults struct {
	Results []TrendingResult `json:"results"`
}

type TrendingResult struct {
	Adult              bool               `json:"adult,omitempty"`
	Gender             int                `json:"gender,omitempty"`
	BackdropPath       string             `json:"backdrop_path,omitempty"`
	GenreIDs           []int64            `json:"genre_ids,omitempty"`
	ID                 int64              `json:"id"`
	OriginalLanguage   string             `json:"original_language"`
	OriginalTitle      string             `json:"original_title,omitempty"`
	Overview           string             `json:"overview,omitempty"`
	PosterPath         string             `json:"poster_path,omitempty"`
	ReleaseDate        string             `json:"release_date,omitempty"`
	Title              string             `json:"title,omitempty"`
	Video              bool               `json:"video,omitempty"`
	VoteAverage        float32            `json:"vote_average,omitempty"`
	VoteCount          int64              `json:"vote_count,omitempty"`
	Popularity         float32            `json:"popularity,omitempty"`
	FirstAirDate       string             `json:"first_air_date,omitempty"`
	Name               string             `json:"name,omitempty"`
	OriginCountry      []string           `json:"origin_country,omitempty"`
	OriginalName       string             `json:"original_name,omitempty"`
	KnownForDepartment string             `json:"known_for_department,omitempty"`
	ProfilePath        string             `json:"profile_path,omitempty"`
	KnownFor           []TrendingKnownFor `json:"known_for,omitempty"`
}

type TrendingKnownFor struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Popularity       float64 `json:"popularity"`
	MediaType        string  `json:"media_type"`
}

// MovieReleaseDatesResults Result Types
type MovieReleaseDatesResults struct {
	Results []MovieReleaseDatesResult `json:"results"`
}

type MovieReleaseDatesResult struct {
	CountryCode  string                   `json:"iso_3166_1"`
	ReleaseDates []MovieReleaseDateResult `json:"release_dates"`
}

type MovieReleaseDateResult struct {
	Certification string `json:"certification"`
	LanguageCode  string `json:"iso_639_1"`
	ReleaseDate   string `json:"release_date"`
	Type          int    `json:"type"`
	Note          string `json:"note"`
}

// MovieVideosResults Result Types
type MovieVideosResults struct {
	Results []MovieVideosResult `json:"results"`
}

type MovieVideosResult struct {
	ID           string `json:"id"`
	LanguageCode string `json:"iso_639_1"`
	CountryCode  string `json:"iso_3166_1"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Official     bool   `json:"official"`
	PublishedAt  string `json:"published_at"`
	Site         string `json:"site"`
	Size         int    `json:"size"`
	Type         string `json:"type"`
}

// MovieWatchProvidersResults Result Types
type MovieWatchProvidersResults struct {
	Results map[string]MovieWatchProvidersResult `json:"results"`
}

type MovieWatchProvidersResult struct {
	Link     string          `json:"link"`
	FlatRate []WatchProvider `json:"flatrate,omitempty"`
	Rent     []WatchProvider `json:"rent,omitempty"`
	Buy      []WatchProvider `json:"buy,omitempty"`
}

// MovieRecommendationsResults Result Types
type MovieRecommendationsResults struct {
	Results []MovieRecommendationsResult `json:"results"`
}

type MovieRecommendationsResult struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	GenreIDs         []int64 `json:"genre_ids"`
	ID               int64   `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float32 `json:"popularity"`
	VoteCount        int64   `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
}

// MovieReviewsResults Result Types
type MovieReviewsResults struct {
	Results []MovieReviewsResult `json:"results"`
}

type MovieReviewsResult struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

// MovieListsResults Result Types
type MovieListsResults struct {
	Results []MovieListsResult `json:"results"`
}

type MovieListsResult struct {
	Description   string `json:"description"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	ItemCount     int64  `json:"item_count"`
	LanguageCode  string `json:"iso_639_1"`
	ListType      string `json:"list_type"`
	Name          string `json:"name"`
	PosterPath    string `json:"poster_path"`
}

// MovieNowPlayingResults Result Types
type MovieNowPlayingResults struct {
	Results []MovieNowPlayingResult `json:"results"`
}

type MovieNowPlayingResult struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	Genres           []Genre `json:"genres"`
	ID               int64   `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float32 `json:"popularity"`
	VoteCount        int64   `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
}

// MoviePopularResults Result Types
type MoviePopularResults struct {
	Results []MoviePopularResult `json:"results"`
}

type MoviePopularResult struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	Genres           []Genre `json:"genres"`
	ID               int64   `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float32 `json:"popularity"`
	VoteCount        int64   `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
}

// TVAlternativeTitlesResults Result Types
type TVAlternativeTitlesResults struct {
	Results []TVAlternativeTitlesResult `json:"results"`
}

type TVAlternativeTitlesResult struct {
	CountryCode string `json:"iso_3166_1"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

// TVContentRatingsResults Result Types
type TVContentRatingsResults struct {
	Results []TVContentRatingsResult `json:"results"`
}

type TVContentRatingsResult struct {
	CountryCode string `json:"iso_3166_1"`
	Rating      string `json:"rating"`
}

// TVEpisodeGroupsResults Result Types
type TVEpisodeGroupsResults struct {
	Results []TVEpisodeGroupsResult `json:"results"`
}

type TVEpisodeGroupsResult struct {
	Description  string                `json:"description"`
	EpisodeCount int                   `json:"episode_count"`
	GroupCount   int                   `json:"group_count"`
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Network      TVEpisodeGroupNetwork `json:"network"`
	Type         int                   `json:"type"`
}

type TVEpisodeGroupNetwork struct {
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// TVKeywordsResults Result Types
type TVKeywordsResults struct {
	Results []TVKeywordsResult `json:"results"`
}

type TVKeywordsResult struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TVRecommendationsResults Result Types
type TVRecommendationsResults struct {
	Results []TVRecommendationsResult `json:"results"`
}

type TVRecommendationsResult struct {
	PosterPath       string   `json:"poster_path"`
	Popularity       float32  `json:"popularity"`
	ID               int64    `json:"id"`
	BackdropPath     string   `json:"backdrop_path"`
	VoteAverage      float32  `json:"vote_average"`
	Overview         string   `json:"overview"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	GenreIDs         []int64  `json:"genre_ids"`
	OriginalLanguage string   `json:"original_language"`
	VoteCount        int64    `json:"vote_count"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
}

// TVReviewsResults Result Types
type TVReviewsResults struct {
	Results []TVReviewsResult `json:"results"`
}

type TVReviewsResult struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

// TVScreenedTheatricallyResults Result Types
type TVScreenedTheatricallyResults struct {
	Results []TVScreenedTheatricallyResult `json:"results"`
}

type TVScreenedTheatricallyResult struct {
	ID            int64 `json:"id"`
	EpisodeNumber int   `json:"episode_number"`
	SeasonNumber  int   `json:"season_number"`
}

// TVWatchProvidersResults Result Types
type TVWatchProvidersResults struct {
	Results map[string]TVWatchProvidersResult `json:"results"`
}

type TVWatchProvidersResult struct {
	Link     string          `json:"link"`
	FlatRate []WatchProvider `json:"flatrate,omitempty"`
	Rent     []WatchProvider `json:"rent,omitempty"`
	Buy      []WatchProvider `json:"buy,omitempty"`
}

// TVVideosResults Result Types
type TVVideosResults struct {
	Results []TVVideosResult `json:"results"`
}

type TVVideosResult struct {
	ID           string `json:"id"`
	LanguageCode string `json:"iso_639_1"`
	CountryCode  string `json:"iso_3166_1"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Site         string `json:"site"`
	Size         int    `json:"size"`
	Type         string `json:"type"`
}

// TVAiringTodayResults Result Types
type TVAiringTodayResults struct {
	Results []TVAiringTodayResult `json:"results"`
}

type TVAiringTodayResult struct {
	OriginalName     string   `json:"original_name"`
	GenreIDs         []int64  `json:"genre_ids"`
	Name             string   `json:"name"`
	Popularity       float32  `json:"popularity"`
	OriginCountry    []string `json:"origin_country"`
	VoteCount        int64    `json:"vote_count"`
	FirstAirDate     string   `json:"first_air_date"`
	BackdropPath     string   `json:"backdrop_path"`
	OriginalLanguage string   `json:"original_language"`
	ID               int64    `json:"id"`
	VoteAverage      float32  `json:"vote_average"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
}
