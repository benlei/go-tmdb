package tmdb

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// TVEpisodeGroupsDetails type is a struct for details JSON response.
type TVEpisodeGroupsDetails struct {
	Description  string                `json:"description"`
	EpisodeCount int                   `json:"episode_count"`
	GroupCount   int                   `json:"group_count"`
	Groups       []TVEpisodeGroup      `json:"groups"`
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Network      TVEpisodeGroupNetwork `json:"network"`
	Type         int                   `json:"type"`
}

type TVEpisodeGroup struct {
	ID       string                  `json:"id"`
	Name     string                  `json:"name"`
	Order    int                     `json:"order"`
	Episodes []TVEpisodeGroupEpisode `json:"episodes"`
	Locked   bool                    `json:"locked"`
}

type TVEpisodeGroupEpisode struct {
	AirDate        string              `json:"air_date"`
	EpisodeNumber  int                 `json:"episode_number"`
	ID             int64               `json:"id"`
	Name           string              `json:"name"`
	Overview       string              `json:"overview"`
	ProductionCode jsoniter.RawMessage `json:"production_code"`
	SeasonNumber   int                 `json:"season_number"`
	ShowID         int64               `json:"show_id"`
	StillPath      string              `json:"still_path"`
	VoteAverage    float32             `json:"vote_average"`
	VoteCount      int64               `json:"vote_count"`
	Order          int                 `json:"order"`
}

// GetTVEpisodeGroupsDetails the details of a TV episode group.
// Groups support 7 different types which are enumerated as the following:
//
// 1. Original air date
// 2. Absolute
// 3. DVD
// 4. Digital
// 5. Story arc
// 6. Production
// 7. TV
//
// https://developers.themoviedb.org/3/tv-episode-groups/get-tv-episode-group-details
func (c *Client) GetTVEpisodeGroupsDetails(
	id string,
	urlOptions map[string]string,
) (*TVEpisodeGroupsDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sepisode_group/%s?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvEpisodeGroupDetails := TVEpisodeGroupsDetails{}
	if err := c.get(tmdbURL, &tvEpisodeGroupDetails); err != nil {
		return nil, err
	}
	return &tvEpisodeGroupDetails, nil
}
