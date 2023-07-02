package tmdb

import "fmt"

// TVSeasonDetails is a struct for details JSON response.
type TVSeasonDetails struct {
	IDString     string             `json:"_id"`
	AirDate      string             `json:"air_date"`
	Episodes     []*TVSeasonEpisode `json:"episodes"`
	Name         string             `json:"name"`
	Overview     string             `json:"overview"`
	ID           int64              `json:"id"`
	PosterPath   string             `json:"poster_path"`
	SeasonNumber int                `json:"season_number"`
	*TVSeasonCreditsAppend
	*TVSeasonExternalIDsAppend
	*TVSeasonImagesAppend
	*TVSeasonVideosAppend
}

type TVSeasonEpisode struct {
	ShowID int64 `json:"show_id"`
	*TVEpisodeDetails
}

// TVSeasonCreditsAppend type is a struct
// for credits in append to response.
type TVSeasonCreditsAppend struct {
	Credits *TVSeasonCredits `json:"credits,omitempty"`
}

// TVSeasonExternalIDsAppend type is a struct
// for external ids in append to response.
type TVSeasonExternalIDsAppend struct {
	ExternalIDs *TVSeasonExternalIDs `json:"external_ids,omitempty"`
}

// TVSeasonImagesAppend type is a struct
// for images in append to response.
type TVSeasonImagesAppend struct {
	Images *TVSeasonImages `json:"images,omitempty"`
}

// TVSeasonVideosAppend type is a struct
// for videos in append to response.
type TVSeasonVideosAppend struct {
	Videos *TVSeasonVideos `json:"videos,omitempty"`
}

// TVSeasonChanges is a struct for changes JSON response.
type TVSeasonChanges ChangeSet

// TVSeasonCredits type is a struct for credits JSON response.
type TVSeasonCredits struct {
	Cast []*TVSeasonCreditCast `json:"cast"`
	Crew []*TVSeasonCreditCrew `json:"crew"`
	ID   int                   `json:"id"`
}

type TVSeasonCreditCast struct {
	Character   string `json:"character"`
	CreditID    string `json:"credit_id"`
	Gender      int    `json:"gender"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	ProfilePath string `json:"profile_path"`
}

type TVSeasonCreditCrew struct {
	CreditID    string `json:"credit_id"`
	Department  string `json:"department"`
	Gender      int    `json:"gender"`
	ID          int64  `json:"id"`
	Job         string `json:"job"`
	Name        string `json:"name"`
	ProfilePath string `json:"profile_path"`
}

// TVSeasonExternalIDs type is a struct for external ids JSON response.
type TVSeasonExternalIDs struct {
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	ID          int64  `json:"id,omitempty"`
}

// TVSeasonImages type is a struct for images JSON response.
type TVSeasonImages struct {
	ID      int64    `json:"id,omitempty"`
	Posters []*Image `json:"posters"`
}

// TVSeasonVideos type is a struct for videos JSON response.
type TVSeasonVideos struct {
	ID      int64                  `json:"id,omitempty"`
	Results []*TVSeasonVideoResult `json:"results"`
}

type TVSeasonVideoResult struct {
	ID           string `json:"id"`
	LanguageCode string `json:"iso_639_1"`
	CountryCode  string `json:"iso_3166_1"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Site         string `json:"site"`
	Size         int    `json:"size"`
	Type         string `json:"type"`
}

// GetTVSeasonDetails get the TV season details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-details
func (c *Client) GetTVSeasonDetails(id int64, seasonNumber int, urlOptions map[string]string) (*TVSeasonDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonDetails := TVSeasonDetails{}
	if err := c.get(tmdbURL, &tvSeasonDetails); err != nil {
		return nil, err
	}
	return &tvSeasonDetails, nil
}

// GetTVSeasonChanges get the changes for a TV season.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-changes
func (c *Client) GetTVSeasonChanges(id int64, urlOptions map[string]string) (*TVSeasonChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sseason/%d/changes?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvSeasonChanges := TVSeasonChanges{}
	if err := c.get(tmdbURL, &tvSeasonChanges); err != nil {
		return nil, err
	}
	return &tvSeasonChanges, nil
}

// GetTVSeasonCredits get the credits for TV season.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-credits
func (c *Client) GetTVSeasonCredits(id int64, seasonNumber int, urlOptions map[string]string) (*TVSeasonCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/credits?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tVSeasonCredits := TVSeasonCredits{}
	if err := c.get(tmdbURL, &tVSeasonCredits); err != nil {
		return nil, err
	}
	return &tVSeasonCredits, nil
}

// GetTVSeasonExternalIDs get the external ids for a TV season.
// We currently support the following external sources.
//
// Media Databases: TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// *Defunct or no longer available as a service.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-external-ids
func (c *Client) GetTVSeasonExternalIDs(id int64, seasonNumber int, urlOptions map[string]string) (*TVSeasonExternalIDs, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/external_ids?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonExternalIDs := TVSeasonExternalIDs{}
	if err := c.get(tmdbURL, &tvSeasonExternalIDs); err != nil {
		return nil, err
	}
	return &tvSeasonExternalIDs, nil
}

// GetTVSeasonImages get the images that belong to a TV season.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-images
func (c *Client) GetTVSeasonImages(id int64, seasonNumber int, urlOptions map[string]string) (*TVSeasonImages, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/images?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonImages := TVSeasonImages{}
	if err := c.get(tmdbURL, &tvSeasonImages); err != nil {
		return nil, err
	}
	return &tvSeasonImages, nil
}

// GetTVSeasonVideos get the videos that have been added to a TV season.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-videos
func (c *Client) GetTVSeasonVideos(id int64, seasonNumber int, urlOptions map[string]string) (*TVSeasonVideos, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/videos?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonVideos := TVSeasonVideos{}
	if err := c.get(tmdbURL, &tvSeasonVideos); err != nil {
		return nil, err
	}
	return &tvSeasonVideos, nil
}
