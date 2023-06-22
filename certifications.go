package tmdb

import "fmt"

// Certification type is a struct for a single certification JSON response.
type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int    `json:"order"`
}

type MovieCertifications struct {
	AU    []*Certification `json:"AU"`
	BG    []*Certification `json:"BG"`
	BR    []*Certification `json:"BR"`
	CA    []*Certification `json:"CA"`
	CA_QC []*Certification `json:"CA-QC"`
	DE    []*Certification `json:"DE"`
	DK    []*Certification `json:"DK"`
	ES    []*Certification `json:"ES"`
	FI    []*Certification `json:"FI"`
	FR    []*Certification `json:"FR"`
	GB    []*Certification `json:"GB"`
	HU    []*Certification `json:"HU"`
	IN    []*Certification `json:"IN"`
	IT    []*Certification `json:"IT"`
	LT    []*Certification `json:"LT"`
	MY    []*Certification `json:"MY"`
	NL    []*Certification `json:"NL"`
	NO    []*Certification `json:"NO"`
	NZ    []*Certification `json:"NZ"`
	PH    []*Certification `json:"PH"`
	PT    []*Certification `json:"PT"`
	RU    []*Certification `json:"RU"`
	SE    []*Certification `json:"SE"`
	US    []*Certification `json:"US"`
}

// CertificationMovie type is a struct for movie certifications JSON response.
type CertificationMovie struct {
	Certifications MovieCertifications `json:"certifications"`
}

type TVCertifications struct {
	AU    []*Certification `json:"AU"`
	BR    []*Certification `json:"BR"`
	CA    []*Certification `json:"CA"`
	CA_QC []*Certification `json:"CA-QC"`
	DE    []*Certification `json:"DE"`
	ES    []*Certification `json:"ES"`
	FR    []*Certification `json:"FR"`
	GB    []*Certification `json:"GB"`
	HU    []*Certification `json:"HU"`
	KR    []*Certification `json:"KR"`
	LT    []*Certification `json:"LT"`
	NL    []*Certification `json:"NL"`
	PH    []*Certification `json:"PH"`
	PT    []*Certification `json:"PT"`
	RU    []*Certification `json:"RU"`
	SK    []*Certification `json:"SK"`
	TH    []*Certification `json:"TH"`
	US    []*Certification `json:"US"`
}

// CertificationTV type is a struct for tv certifications JSON response.
type CertificationTV struct {
	Certifications TVCertifications `json:"certifications"`
}

// GetCertificationMovie get an up to date list of the
// officially supported movie certifications on TMDb.
//
// https://developers.themoviedb.org/3/certifications/get-movie-certifications
func (c *Client) GetCertificationMovie() (
	*CertificationMovie,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL,
		movieURL,
		c.apiKey,
	)
	certificationMovie := CertificationMovie{}
	if err := c.get(tmdbURL, &certificationMovie); err != nil {
		return nil, err
	}
	return &certificationMovie, nil
}

// GetCertificationTV get an up to date list of the
// officially supported TV show certifications on TMDb.
//
// https://developers.themoviedb.org/3/certifications/get-tv-certifications
func (c *Client) GetCertificationTV() (
	*CertificationTV,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL,
		tvURL,
		c.apiKey,
	)
	certificationTV := CertificationTV{}
	if err := c.get(tmdbURL, &certificationTV); err != nil {
		return nil, err
	}
	return &certificationTV, nil
}
