package tmdb

import (
	"fmt"
)

// TVEpisodeDetails type is a struct for details JSON response.
type TVEpisodeDetails struct {
	AirDate        string                `json:"air_date"`
	EpisodeNumber  int                   `json:"episode_number"`
	Crew           []*TVEpisodeCrew      `json:"crew"`
	GuestStars     []*TVEpisodeGuestStar `json:"guest_stars"`
	Name           string                `json:"name"`
	Overview       string                `json:"overview"`
	ID             int64                 `json:"id"`
	ProductionCode string                `json:"production_code"`
	SeasonNumber   int                   `json:"season_number"`
	StillPath      string                `json:"still_path"`
	VoteAverage    float32               `json:"vote_average"`
	VoteCount      int64                 `json:"vote_count"`
	*TVEpisodeCreditsAppend
	*TVEpisodeExternalIDsAppend
	*TVEpisodeImagesAppend
	*TVEpisodeTranslationsAppend
	*TVEpisodeVideosAppend
}

type TVEpisodeCrew struct {
	ID          int64  `json:"id"`
	CreditID    string `json:"credit_id"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Job         string `json:"job"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type TVEpisodeGuestStar struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CreditID    string `json:"credit_id"`
	Character   string `json:"character"`
	Order       int    `json:"order"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

// TVEpisodeCreditsAppend type is a struct
// for credits in append to response.
type TVEpisodeCreditsAppend struct {
	Credits *TVEpisodeCredits `json:"credits,omitempty"`
}

// TVEpisodeExternalIDsAppend type is a struct
// for external ids in append to response.
type TVEpisodeExternalIDsAppend struct {
	ExternalIDs *TVEpisodeExternalIDs `json:"external_ids,omitempty"`
}

// TVEpisodeImagesAppend type is a struct
// for images in append to response.
type TVEpisodeImagesAppend struct {
	Images *TVEpisodeImages `json:"images,omitempty"`
}

// TVEpisodeTranslationsAppend type is a struct
// for translations in append to response.
type TVEpisodeTranslationsAppend struct {
	Translations *TVEpisodeTranslations `json:"translations,omitempty"`
}

// TVEpisodeVideosAppend type is a struct
// for videos in append to response.
type TVEpisodeVideosAppend struct {
	Videos *TVEpisodeVideos `json:"videos,omitempty"`
}

// TVEpisodeChanges type is a struct for changes JSON response.
type TVEpisodeChanges ChangeSet

// TVEpisodeCredits type is a struct for credits JSON response.
type TVEpisodeCredits struct {
	Cast       []*TVEpisodeCreditCast      `json:"cast"`
	Crew       []*TVEpisodeCreditCrew      `json:"crew"`
	GuestStars []*TVEpisodeCreditGuestStar `json:"guest_stars"`
	ID         int64                       `json:"id,omitempty"`
}

type TVEpisodeCreditCast struct {
	Character   string `json:"character"`
	CreditID    string `json:"credit_id"`
	Gender      int    `json:"gender"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	ProfilePath string `json:"profile_path"`
}

type TVEpisodeCreditCrew struct {
	ID          int64  `json:"id"`
	CreditID    string `json:"credit_id"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Job         string `json:"job"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type TVEpisodeCreditGuestStar struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CreditID    string `json:"credit_id"`
	Character   string `json:"character"`
	Order       int    `json:"order"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

// TVEpisodeExternalIDs type is a struct for external ids JSON response.
type TVEpisodeExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	IMDbID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
}

// TVEpisodeImages type is a struct for images JSON response.
type TVEpisodeImages struct {
	ID     int64    `json:"id,omitempty"`
	Stills []*Image `json:"stills"`
}

// TVEpisodeTranslations type is a struct for translations JSON response.
type TVEpisodeTranslations struct {
	ID           int64                   `json:"id,omitempty"`
	Translations []*TVEpisodeTranslation `json:"translations"`
}

type TVEpisodeTranslation struct {
	CountryCode  string                    `json:"iso_3166_1"`
	LanguageCode string                    `json:"iso_639_1"`
	Name         string                    `json:"name"`
	EnglishName  string                    `json:"english_name"`
	Data         *TVEpisodeTranslationData `json:"data"`
}

type TVEpisodeTranslationData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

// TVEpisodeRate type is a struct for rate JSON response.
type TVEpisodeRate struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// TVEpisodeVideos type is a struct for videos JSON response.
type TVEpisodeVideos struct {
	ID      int64                   `json:"id,omitempty"`
	Results []*TVEpisodeVideoResult `json:"results"`
}

type TVEpisodeVideoResult struct {
	ID           string `json:"id"`
	LanguageCode string `json:"iso_639_1"`
	CountryCode  string `json:"iso_3166_1"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Site         string `json:"site"`
	Size         int    `json:"size"`
	Type         string `json:"type"`
}

// GetTVEpisodeDetails get the TV episode details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-details
func (c *Client) GetTVEpisodeDetails(id int64, seasonNumber int, episodeNumber int, urlOptions map[string]string) (*TVEpisodeDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
		options,
	)
	tvEpisodeDetails := TVEpisodeDetails{}
	if err := c.get(tmdbURL, &tvEpisodeDetails); err != nil {
		return nil, err
	}
	return &tvEpisodeDetails, nil
}

// GetTVEpisodeChanges get the changes for a TV episode.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-changes
func (c *Client) GetTVEpisodeChanges(id int64, urlOptions map[string]string) (*TVEpisodeChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sepisode/%d/changes?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvEpisodeChanges := TVEpisodeChanges{}
	if err := c.get(tmdbURL, &tvEpisodeChanges); err != nil {
		return nil, err
	}
	return &tvEpisodeChanges, nil
}

// GetTVEpisodeCredits get the credits (cast, crew and guest stars) for a TV episode.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-credits
func (c *Client) GetTVEpisodeCredits(id int64, seasonNumber int, episodeNumber int) (*TVEpisodeCredits, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/credits?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeCredits := TVEpisodeCredits{}
	if err := c.get(tmdbURL, &tvEpisodeCredits); err != nil {
		return nil, err
	}
	return &tvEpisodeCredits, nil
}

// GetTVEpisodeExternalIDs get the external ids for a TV episode.
// We currently support the following external sources.
//
// Media Databases: IMDb ID, TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// *Defunct or no longer available as a service.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-external-ids
func (c *Client) GetTVEpisodeExternalIDs(id int64, seasonNumber int, episodeNumber int) (*TVEpisodeExternalIDs, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/external_ids?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeExternalIDs := TVEpisodeExternalIDs{}
	if err := c.get(tmdbURL, &tvEpisodeExternalIDs); err != nil {
		return nil, err
	}
	return &tvEpisodeExternalIDs, nil
}

// GetTVEpisodeImages get the images that belong to a TV episode.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-images
func (c *Client) GetTVEpisodeImages(id int64, seasonNumber int, episodeNumber int) (*TVEpisodeImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/images?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeImages := TVEpisodeImages{}
	if err := c.get(tmdbURL, &tvEpisodeImages); err != nil {
		return nil, err
	}
	return &tvEpisodeImages, nil
}

// GetTVEpisodeTranslations get the translation data for an episode.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-translations
func (c *Client) GetTVEpisodeTranslations(id int64, seasonNumber int, episodeNumber int) (*TVEpisodeTranslations, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/translations?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeTranslations := TVEpisodeTranslations{}
	if err := c.get(tmdbURL, &tvEpisodeTranslations); err != nil {
		return nil, err
	}
	return &tvEpisodeTranslations, nil
}

// GetTVEpisodeVideos get the videos that have been added to a TV episode.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-videos
func (c *Client) GetTVEpisodeVideos(id int64, seasonNumber int, episodeNumber int, urlOptions map[string]string) (*TVEpisodeVideos, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/videos?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
		options,
	)
	tvEpisodeVideos := TVEpisodeVideos{}
	if err := c.get(tmdbURL, &tvEpisodeVideos); err != nil {
		return nil, err
	}
	return &tvEpisodeVideos, nil
}
