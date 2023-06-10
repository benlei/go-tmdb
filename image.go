package tmdb

type Image struct {
	ID          string  `json:"id,omitempty"`
	AspectRatio float32 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	Iso639_1    string  `json:"iso_639_1"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	Width       int     `json:"width"`
}
