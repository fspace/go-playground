package ch1

import "fmt"

type Movie struct {
	Actors      []string
	Rating      float32
	ReleaseYear int
	Title       string
}

func (movie Movie) DisplayTitle() string {
	return fmt.Sprintf("%s (%d)", movie.Title, movie.ReleaseYear)
}

// ----------------------------------------------------------------------
type Counter struct {
	Count int
}

func (c Counter) Increment() {
	c.Count++
}
func (c *Counter) IncrementWithPointer() {
	c.Count++
}

// ------------------------------------------------------------------------
// myUnexportedFunc is not available to code outside this package
func myUnexportedFunc() {
}

// MyExportedFunc is available outside this package, though
func MyExportedFunc() {
}

type MyExportedType struct {
	ExportedField   string
	unexportedField string
}

// ==========================================================================================
// API

func StructsCreating() {
	episodeIV := Movie{
		Title:       "Star Wars: A New Hope",
		Rating:      5.0,
		ReleaseYear: 1977,
	}

	episodeIV.Actors = []string{
		"Mark Hamill",
		"Harrison Ford",
		"Carrie Fisher",
	}
	fmt.Println(episodeIV.Title, "has a rating of", episodeIV.Rating)
}

func TypeMethods() {
	episodeV := Movie{
		Title:       "Star Wars: The Empire Strikes Back",
		ReleaseYear: 1980,
	}
	fmt.Println(episodeV.DisplayTitle())
	// Outputs: “Star Wars: The Empire Strikes Back (1980)”
}

func TypeMethods2() {
	counter := &Counter{}
	fmt.Println(counter.Count) // Outputs: 0
	counter.Increment()
	fmt.Println(counter.Count) // Outputs: 0
	counter.IncrementWithPointer()
	fmt.Println(counter.Count) // Outputs: 1
}

// ==========================================================================================
