// Package tmdb initially generated by interfacer
package tmdb

import "net/http"

type TMDBClient interface {
	AddMovie(int64, *ListMedia) (*Response, error)
	AddToWatchlist(int64, *AccountWatchlist) (*Response, error)
	ClearList(int64, bool) (*Response, error)
	CreateGuestSession() (*RequestToken, error)
	CreateList(*ListCreate) (*ListResponse, error)
	CreateRequestToken() (*RequestToken, error)
	DeleteList(int64) (*Response, error)
	DeleteMovieRating(int64, map[string]string) (*Response, error)
	DeleteTVShowRating(int64, map[string]string) (*Response, error)
	GetAccountDetails() (*AccountDetails, error)
	GetAvailableWatchProviderRegions(map[string]string) (*WatchRegionList, error)
	GetCertificationMovie() (*CertificationMovie, error)
	GetCertificationTV() (*CertificationTV, error)
	GetChangesMovie(map[string]string) (*ChangesMovie, error)
	GetChangesPerson(map[string]string) (*ChangesPerson, error)
	GetChangesTV(map[string]string) (*ChangesTV, error)
	GetCollectionDetails(int64, map[string]string) (*CollectionDetails, error)
	GetCollectionImages(int64, map[string]string) (*CollectionImages, error)
	GetCollectionTranslations(int64, map[string]string) (*CollectionTranslations, error)
	GetCompanyAlternativeNames(int64) (*CompanyAlternativeNames, error)
	GetCompanyDetails(int64) (*CompanyDetails, error)
	GetCompanyImages(int64) (*CompanyImages, error)
	GetConfigurationAPI() (*ConfigurationAPI, error)
	GetConfigurationCountries() (*ConfigurationCountries, error)
	GetConfigurationJobs() (*ConfigurationJobs, error)
	GetConfigurationLanguages() (*ConfigurationLanguages, error)
	GetConfigurationPrimaryTranslations() (*ConfigurationPrimaryTranslations, error)
	GetConfigurationTimezones() (*ConfigurationTimezones, error)
	GetCreatedLists(int64, map[string]string) (*AccountCreatedLists, error)
	GetCreditDetails(string) (*CreditsDetails, error)
	GetDiscoverMovie(map[string]string) (*DiscoverMovie, error)
	GetDiscoverTV(map[string]string) (*DiscoverTV, error)
	GetFavoriteMovies(int64, map[string]string) (*AccountFavoriteMovies, error)
	GetFavoriteTVShows(int64, map[string]string) (*AccountFavoriteTVShows, error)
	GetFindByID(string, map[string]string) (*FindByID, error)
	GetGenreMovieList(map[string]string) (*GenreMovieList, error)
	GetGenreTVList(map[string]string) (*GenreMovieList, error)
	GetGuestSessionRatedMovies(string, map[string]string) (*GuestSessionRatedMovies, error)
	GetGuestSessionRatedTVEpisodes(string, map[string]string) (*GuestSessionRatedTVEpisodes, error)
	GetGuestSessionRatedTVShows(string, map[string]string) (*GuestSessionRatedTVShows, error)
	GetKeywordDetails(int64) (*KeywordDetails, error)
	GetKeywordMovies(int64, map[string]string) (*KeywordMovies, error)
	GetListDetails(string, map[string]string) (*ListDetails, error)
	GetListItemStatus(string, map[string]string) (*ListItemStatus, error)
	GetMovieAccountStates(int, map[string]string) (*MovieAccountStates, error)
	GetMovieAlternativeTitles(int, map[string]string) (*MovieAlternativeTitles, error)
	GetMovieChanges(int64, map[string]string) (*MovieChanges, error)
	GetMovieCredits(int64, map[string]string) (*MovieCredits, error)
	GetMovieDetails(int, map[string]string) (*MovieDetails, error)
	GetMovieExternalIDs(int, map[string]string) (*MovieExternalIDs, error)
	GetMovieImages(int, map[string]string) (*MovieImages, error)
	GetMovieKeywords(int) (*MovieKeywords, error)
	GetMovieLatest(map[string]string) (*MovieLatest, error)
	GetMovieLists(int, map[string]string) (*MovieLists, error)
	GetMovieNowPlaying(map[string]string) (*MovieNowPlaying, error)
	GetMoviePopular(map[string]string) (*MoviePopular, error)
	GetMovieRecommendations(int, map[string]string) (*MovieRecommendations, error)
	GetMovieReleaseDates(int) (*MovieReleaseDates, error)
	GetMovieReviews(int, map[string]string) (*MovieReviews, error)
	GetMovieSimilar(int, map[string]string) (*MovieSimilar, error)
	GetMovieTopRated(map[string]string) (*MovieTopRated, error)
	GetMovieTranslations(int, map[string]string) (*MovieTranslations, error)
	GetMovieUpcoming(map[string]string) (*MovieUpcoming, error)
	GetMovieVideos(int, map[string]string) (*MovieVideos, error)
	GetMovieWatchProviders(int, map[string]string) (*MovieWatchProviders, error)
	GetMovieWatchlist(int, map[string]string) (*AccountMovieWatchlist, error)
	GetNetworkAlternativeNames(int) (*NetworkAlternativeNames, error)
	GetNetworkDetails(int) (*NetworkDetails, error)
	GetNetworkImages(int) (*NetworkImages, error)
	GetPersonChanges(int, map[string]string) (*PersonChanges, error)
	GetPersonCombinedCredits(int, map[string]string) (*PersonCombinedCredits, error)
	GetPersonDetails(int, map[string]string) (*PersonDetails, error)
	GetPersonExternalIDs(int, map[string]string) (*PersonExternalIDs, error)
	GetPersonImages(int) (*PersonImages, error)
	GetPersonLatest(map[string]string) (*PersonLatest, error)
	GetPersonMovieCredits(int, map[string]string) (*PersonMovieCredits, error)
	GetPersonPopular(map[string]string) (*PersonPopular, error)
	GetPersonTVCredits(int, map[string]string) (*PersonTVCredits, error)
	GetPersonTaggedImages(int, map[string]string) (*PersonTaggedImages, error)
	GetPersonTranslations(int, map[string]string) (*PersonTranslations, error)
	GetRatedMovies(int, map[string]string) (*AccountRatedMovies, error)
	GetRatedTVEpisodes(int, map[string]string) (*AccountRatedTVEpisodes, error)
	GetRatedTVShows(int, map[string]string) (*AccountRatedTVShows, error)
	GetReviewDetails(string) (*ReviewDetails, error)
	GetSearchCollections(string, map[string]string) (*SearchCollections, error)
	GetSearchCompanies(string, map[string]string) (*SearchCompanies, error)
	GetSearchKeywords(string, map[string]string) (*SearchKeywords, error)
	GetSearchMovies(string, map[string]string) (*SearchMovies, error)
	GetSearchMulti(string, map[string]string) (*SearchMulti, error)
	GetSearchPeople(string, map[string]string) (*SearchPeople, error)
	GetSearchTVShow(string, map[string]string) (*SearchTVShows, error)
	GetTVAccountStates(int, map[string]string) (*TVAccountStates, error)
	GetTVAggregateCredits(int, map[string]string) (*TVAggregateCredits, error)
	GetTVAiringToday(map[string]string) (*TVAiringToday, error)
	GetTVAlternativeTitles(int, map[string]string) (*TVAlternativeTitles, error)
	GetTVChanges(int, map[string]string) (*TVChanges, error)
	GetTVContentRatings(int, map[string]string) (*TVContentRatings, error)
	GetTVCredits(int, map[string]string) (*TVCredits, error)
	GetTVDetails(int, map[string]string) (*TVDetails, error)
	GetTVEpisodeChanges(int, map[string]string) (*TVEpisodeChanges, error)
	GetTVEpisodeCredits(int, int, int) (*TVEpisodeCredits, error)
	GetTVEpisodeDetails(int, int, int, map[string]string) (*TVEpisodeDetails, error)
	GetTVEpisodeExternalIDs(int, int, int) (*TVEpisodeExternalIDs, error)
	GetTVEpisodeGroups(int, map[string]string) (*TVEpisodeGroups, error)
	GetTVEpisodeGroupsDetails(string, map[string]string) (*TVEpisodeGroupsDetails, error)
	GetTVEpisodeImages(int, int, int) (*TVEpisodeImages, error)
	GetTVEpisodeTranslations(int, int, int) (*TVEpisodeTranslations, error)
	GetTVEpisodeVideos(int, int, int, map[string]string) (*TVEpisodeVideos, error)
	GetTVExternalIDs(int, map[string]string) (*TVExternalIDs, error)
	GetTVImages(int, map[string]string) (*TVImages, error)
	GetTVKeywords(int) (*TVKeywords, error)
	GetTVLatest(map[string]string) (*TVLatest, error)
	GetTVOnTheAir(map[string]string) (*TVOnTheAir, error)
	GetTVPopular(map[string]string) (*TVPopular, error)
	GetTVRecommendations(int, map[string]string) (*TVRecommendations, error)
	GetTVReviews(int, map[string]string) (*TVReviews, error)
	GetTVScreenedTheatrically(int) (*TVScreenedTheatrically, error)
	GetTVSeasonChanges(int, map[string]string) (*TVSeasonChanges, error)
	GetTVSeasonCredits(int, int, map[string]string) (*TVSeasonCredits, error)
	GetTVSeasonDetails(int, int, map[string]string) (*TVSeasonDetails, error)
	GetTVSeasonExternalIDs(int, int, map[string]string) (*TVSeasonExternalIDs, error)
	GetTVSeasonImages(int, int, map[string]string) (*TVSeasonImages, error)
	GetTVSeasonVideos(int, int, map[string]string) (*TVSeasonVideos, error)
	GetTVShowsWatchlist(int, map[string]string) (*AccountTVShowsWatchlist, error)
	GetTVSimilar(int, map[string]string) (*TVSimilar, error)
	GetTVTopRated(map[string]string) (*TVTopRated, error)
	GetTVTranslations(int, map[string]string) (*TVTranslations, error)
	GetTVVideos(int, map[string]string) (*TVVideos, error)
	GetTVWatchProviders(int, map[string]string) (*TVWatchProviders, error)
	GetTrending(string, string) (*Trending, error)
	GetWatchProvidersMovie(map[string]string) (*WatchProviderList, error)
	GetWatchProvidersTv(map[string]string) (*WatchProviderList, error)
	MarkAsFavorite(int, *AccountFavorite) (*Response, error)
	PostMovieRating(int, float32, map[string]string) (*Response, error)
	PostTVShowRating(int, float32, map[string]string) (*Response, error)
	RemoveMovie(int, *ListMedia) (*Response, error)
	SetClientAutoRetry()
	SetClientConfig(http.Client)
	SetSessionID(string) error
}

var _ TMDBClient = &Client{}
