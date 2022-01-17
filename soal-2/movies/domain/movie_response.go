package domain

import (
	pb "github.com/muhfaris/stockbit-test/soal-2/movies/grpc/gen/proto"
)

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

// ToProto is cast response to proto model
func (movie MovieResponseModel) ToProto() *pb.MovieResponse {
	var ratings []*pb.Rating
	for _, rating := range movie.Ratings {
		r := pb.Rating{
			Source: rating.Source,
		}

		ratings = append(ratings, &r)
	}

	m := pb.MovieResponse{
		Title:      movie.Title,
		Year:       movie.Year,
		Rated:      movie.Rated,
		Released:   movie.Released,
		Runtime:    movie.Runtime,
		Genre:      movie.Genre,
		Director:   movie.Director,
		Writer:     movie.Writer,
		Actors:     movie.Actors,
		Plot:       movie.Plot,
		Language:   movie.Language,
		Country:    movie.Country,
		Awards:     movie.Awards,
		Poster:     movie.Poster,
		Rating:     ratings,
		Metascore:  movie.Metascore,
		ImdbRating: movie.ImdbRating,
		ImdbVotes:  movie.ImdbVotes,
		ImdbID:     movie.ImdbID,
		Type:       movie.Type,
		DVD:        movie.DVD,
		BoxOffice:  movie.BoxOffice,
		Production: movie.Production,
		Website:    movie.Website,
		Response:   movie.Response,
	}

	return &m
}

// MoviesResponseModel is multiple movie response
type MoviesResponseModel []MovieResponseModel

func (movies MoviesResponseModel) ToProto() []*pb.MovieResponse {
	var data []*pb.MovieResponse
	for _, movie := range movies {

		var ratings []*pb.Rating
		for _, rating := range movie.Ratings {
			r := pb.Rating{
				Source: rating.Source,
			}

			ratings = append(ratings, &r)
		}

		m := pb.MovieResponse{
			Title:      movie.Title,
			Year:       movie.Year,
			Rated:      movie.Rated,
			Released:   movie.Released,
			Runtime:    movie.Runtime,
			Genre:      movie.Genre,
			Director:   movie.Director,
			Writer:     movie.Writer,
			Actors:     movie.Actors,
			Plot:       movie.Plot,
			Language:   movie.Language,
			Country:    movie.Country,
			Awards:     movie.Awards,
			Poster:     movie.Poster,
			Rating:     ratings,
			Metascore:  movie.Metascore,
			ImdbRating: movie.ImdbRating,
			ImdbVotes:  movie.ImdbVotes,
			ImdbID:     movie.ImdbID,
			Type:       movie.Type,
			DVD:        movie.DVD,
			BoxOffice:  movie.BoxOffice,
			Production: movie.Production,
			Website:    movie.Website,
			Response:   movie.Response,
		}

		data = append(data, &m)
	}

	return data
}
