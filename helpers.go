package tmdb

const imageURL = "https://image.tmdb.org/t/p/"

type ImageSize string

func (i ImageSize) String() string {
	return string(i)
}

const (
	// W45 size
	W45 ImageSize = "w45"
	// W92 size
	W92 ImageSize = "w92"
	// W154 size
	W154 ImageSize = "w154"
	// W185 size
	W185 ImageSize = "w185"
	// W300 size
	W300 ImageSize = "w300"
	// W342 size
	W342 ImageSize = "w342"
	// W500 size
	W500 ImageSize = "w500"
	// W780 size
	W780 ImageSize = "w780"
	// W1280 size
	W1280 ImageSize = "w1280"
	// H632 size
	H632 ImageSize = "h632"
	// Original size
	Original ImageSize = "original"
)

// All supported sizes including logos, posters, backdrops,
// profiles and still sizes.
var imageSize = map[ImageSize]string{
	W45:      imageURL + W45.String(),
	W92:      imageURL + W92.String(),
	W154:     imageURL + W154.String(),
	W185:     imageURL + W185.String(),
	W300:     imageURL + W300.String(),
	W342:     imageURL + W342.String(),
	W500:     imageURL + W500.String(),
	W780:     imageURL + W780.String(),
	W1280:    imageURL + W1280.String(),
	H632:     imageURL + H632.String(),
	Original: imageURL + Original.String(),
}

// GetImageURL accepts two parameters, the key and the size
// and returns the complete URL of the image.
//
// Available sizes:
//
// w45      - logo/profile
// w92      - logo/poster/still
// w154     - logo/poster
// w185     - logo/poster/profile/still
// w300     - backdrop/logo/still
// w342     - poster
// w500     - logo/poster
// w780     - backdrop/poster
// w1280    - backdrop
// h632     - profile
// original - backdrop/logo/poster/profile/still
//
// https://developers.themoviedb.org/3/configuration/get-api-configuration
func GetImageURL(key string, size ImageSize) string {
	return imageSize[size] + key
}

const videoURL = "https://www.youtube.com/watch?v="

// GetVideoURL accepts one parameter, the key and returns
// the complete URL of the video.
func GetVideoURL(key string) string {
	return videoURL + key
}
