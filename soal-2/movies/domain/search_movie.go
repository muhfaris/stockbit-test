package domain

// SearchMovieModel is search movie
type SearchMovieModel struct {
	Search MoviesModel `json:"Search"`
}

// Response is wrap response
func (m SearchMovieModel) Response() MoviesResponseModel {
	var data MoviesResponseModel
	for _, movie := range m.Search {
		data = append(data, movie.ToResponse())
	}

	return data
}
