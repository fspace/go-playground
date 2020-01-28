package ch2

type Movie struct {
	Actors      []string
	Rating      float32
	ReleaseYear int
	Title       string
}

func NewMovie(title string, year int) Movie {
	return Movie{
		Title:       title,
		ReleaseYear: year,
		Actors:      []string{},
	}
}
