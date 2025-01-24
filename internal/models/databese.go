package models

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func LoadMovies(filename string) ([]*Movie, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read and skip header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}
	var movies []*Movie
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

		// Parse record
		year, _ := strconv.Atoi(record[3])
		rating, _ := strconv.ParseFloat(record[4], 64)
		genres := strings.Split(record[1], ",")

		movie, err := NewMovie(
			record[0],
			genres,
			record[2],
			year,
			rating,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
