package main

import (
	"encoding/gob"
	"strings"
)

type Character struct {
	Name        string `json:"name" tag:"foo"`
	Surname     string `json:"surname"`
	Job         string `json:"job,omitempty"`
	YearOfBirth int    `json:"year_of_birth,omitempty"`
}

func main() {
	var char = Character{
		Name:    "Albert",
		Surname: "Wilmarth",
		Job:     "assistant professor",
	}
	s := strings.Builder{}
	e := gob.NewEncoder(&s)
	if err := e.Encode(char); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%q", s.String())
}
