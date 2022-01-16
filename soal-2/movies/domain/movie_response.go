package domain

// MovieResponseModel is response movie
type MovieResponseModel struct {
	Title    string `json:"title,omitempty"`
	Year     string `json:"year,omitempty"`
	Rated    string `json:"rated,omitempty"`
	Released string `json:"released,omitempty"`
	Runtime  string `json:"runtime,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Director string `json:"director,omitempty"`
	Writer   string `json:"writer,omitempty"`
	Actors   string `json:"actors,omitempty"`
	Plot     string `json:"plot,omitempty"`
	Language string `json:"language,omitempty"`
	Country  string `json:"country,omitempty"`
	Awards   string `json:"awards,omitempty"`
	Poster   string `json:"poster,omitempty"`
	Ratings  []struct {
		Source string `json:"source,omitempty"`
		Value  string `json:"value,omitempty"`
	} `json:"ratings,omitempty"`
	Metascore  string `json:"metascore,omitempty"`
	ImdbRating string `json:"imdb_rating,omitempty"`
	ImdbVotes  string `json:"imdb_votes,omitempty"`
	ImdbID     string `json:"imdb_id,omitempty"`
	Type       string `json:"type,omitempty"`
	DVD        string `json:"dvd,omitempty"`
	BoxOffice  string `json:"box_office,omitempty"`
	Production string `json:"production,omitempty"`
	Website    string `json:"website,omitempty"`
	Response   string `json:"response,omitempty"`
}

// MoviesResponseModel is multiple movie response
type MoviesResponseModel []MovieResponseModel
