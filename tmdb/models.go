package tmdb

type Movie struct {
	BackdropPath     string  `json:"backdrop_path"`
	ID               int     `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	MediaType        string  `json:"media_type"`
	Adult            bool    `json:"adult"`
	Title            string  `json:"title"`
	OriginalLanguage string  `json:"original_language"`
	Genres			[]Genres `json:"genres"`
	GenreIds         []int   `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
	ReleaseDate      string  `json:"release_date"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
} 
