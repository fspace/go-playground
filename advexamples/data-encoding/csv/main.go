package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := csv.NewReader(strings.NewReader("a,b,c\ne,f,g\n1,2,3"))
	for {
		r, err := r.Read()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(r)
	}
}

func ReadAll() {
	r := csv.NewReader(strings.NewReader("a,b,c\ne,f,g\n1,2,3"))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range records {
		log.Println(r)
	}
}

func Encode() {
	const million = 1000000
	type Country struct {
		Code, Name string
		Population int
	}
	records := []Country{
		{Code: "IT", Name: "Italy", Population: 60 * million},
		{Code: "ES", Name: "Spain", Population: 46 * million},
		{Code: "JP", Name: "Japan", Population: 126 * million},
		{Code: "US", Name: "United States of America", Population: 327 * million},
	}
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	for _, r := range records {
		if err := w.Write([]string{r.Code, r.Name, strconv.Itoa(r.Population)}); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}
	// w.WriteAll            也有一次性写方法
	// The main difference between Write and WriteAll is that the second operation uses more resources and it requires
	// us to convert the records into a slice of string slices before calling it.
}

func customOptions() {
	r := csv.NewReader(strings.NewReader("a b\ne f g\n1"))
	r.Comma = ' '
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range records {
		log.Println(r)
	}
}
