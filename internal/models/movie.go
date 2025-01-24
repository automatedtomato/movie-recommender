package models

import (
	"errors"
	"fmt"
	"math"
)

type Genre string

const (
	GenreAction      Genre = "Action"
	GenreComedy      Genre = "Comedy"
	GenreDrama       Genre = "Drama"
	GenreSciFi       Genre = "Sci-Fi"
	GenreHorror      Genre = "Horror"
	GenreRomance     Genre = "Romance"
	GenreAnimation   Genre = "Animation"
	GenreDocumentary Genre = "Documentary"
)

func IsValidGenre(g Genre) bool {
	validGenres := []Genre{
		GenreAction, GenreComedy, GenreDrama,
		GenreSciFi, GenreHorror, GenreRomance,
		GenreAnimation, GenreDocumentary}

	for _, validGenre := range validGenres {
		if g == validGenre {
			return true
		}
	}
	return false
}

type Movie struct {
	Title       string
	Genres      []string
	Description string
	ReleaseYear int
	Rating      float64
}

// Constructor
func NewMovie(title string, genres []string, description string, releaseYear int, rating float64) (*Movie, error) {

	// Validate input
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if releaseYear < 1888 || releaseYear > 2025 {
		return nil, fmt.Errorf("release year must be between 1888 and 2025, got %d", releaseYear)
	}

	if rating < 0 || rating > 10 {
		return nil, fmt.Errorf("rating must be between 0 and 10, got %.2f", rating)
	}

	if len(genres) == 0 {
		return nil, errors.New("movie must have at least one genre")
	}

	// Check if genres are valid
	for _, genre := range genres {
		if !IsValidGenre(Genre(genre)) {
			return nil, fmt.Errorf("invalid genre: %s", genre)
		}
	}

	return &Movie{
		Title:       title,
		Genres:      genres,
		Description: description,
		ReleaseYear: releaseYear,
		Rating:      rating,
	}, nil
}

func (m Movie) GetDescription() string {
	return m.Description
}

func (m Movie) String() string {
	return fmt.Sprintf("Title: %s\nGenres: %s\nDescription: %s\nReleaseYear: %d\nRating: %f", m.Title, m.Genres, m.Description, m.ReleaseYear, m.Rating)
}

func GetAllGenres() []Genre {
	return []Genre{
		GenreAction, GenreComedy, GenreDrama,
		GenreSciFi, GenreHorror, GenreRomance,
		GenreAnimation, GenreDocumentary}
}

// Converts movie genres to a vector
func (m Movie) ToGenreVector() []int {
	// Get all genres
	allGenres := GetAllGenres()

	// Create a slice with the length of all genres
	vector := make([]int, len(allGenres))

	// Search for each genre
	for _, movieGenre := range m.Genres {
		// Check if genre is in all genres
		for i, genre := range allGenres {
			// If genre is found, set "1" in the location "i"
			if movieGenre == string(genre) {
				vector[i] = 1
				break
			}
		}
	}
	return vector
}

// Calculates the dot product of two vectors
func DotProduct(a, b []int) int {
	sum := 0
	for i := range a {
		sum += a[i] * b[i] // ベクトルの内積（ドット積を求める）
	}
	return sum
}

// Calculates the magnitude (length) of a vector
func Magnitude(v []int) float64 {
	sum := 0
	for _, val := range v {
		sum += val * val // ベクトルの大きさの二乗を求める
	}
	return math.Sqrt(float64(sum)) // 大きさの二乗の平方根　＝ ベクトルのノルム
}

// Calculates the cosine similarity between two movies
func (m Movie) GenreSimilarity(other Movie) float64 {
	v1 := m.ToGenreVector()     // Get vector of current movie
	v2 := other.ToGenreVector() // Get vector of other movie

	dot := DotProduct(v1, v2)
	mag1 := Magnitude(v1)
	mag2 := Magnitude(v2)

	if mag1 == 0 || mag2 == 0 {
		return 0
	}

	return float64(dot) / (mag1 * mag2)
}
