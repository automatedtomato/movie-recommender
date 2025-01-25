package main

import (
	"fmt"
	"movie-recommender/internal/models"
	"movie-recommender/internal/similarity"
)

func main() {
	fmt.Println("\nMovie recommender system starting...")

	// Load movies from CSV
	movies, err := models.LoadMovies("data/movies.csv")
	if err != nil {
		fmt.Printf("Error loading movies: %v\n", err)
		return
	}

	// Print all movies and find some interesting similarities
	fmt.Printf("\nLoaded %d movies\n", len(movies))

	// Let's analyze "The Matrix" similarities
	matrix := movies[0] // First movie in dataset
	fmt.Printf("\nCosine Similarities with %s:\n", matrix.Title)
	for _, movie := range movies[1:] { // Skip comparing with itself
		similarity := matrix.GenreSimilarity(*movie) // Calculate similarity
		if similarity > 0.4 {                        // Show only relatively similar movies
			fmt.Printf("- %s: %.2f\n", movie.Title, similarity)
		}
	}

	// Calculate TF-IDF
	tfidfVectors := similarity.CalculateTFIDF(movies)

	// Find similar movies
	fmt.Printf("\nTF-IDF Similarities with %s:\n", matrix.Title)

	for _, movie := range movies[1:] {
		similarity := similarity.CosineSimilarity(
			tfidfVectors[matrix.Title],
			tfidfVectors[movie.Title],
		)

		if similarity > 0 {
			fmt.Printf("- %s: %.2f\n", movie.Title, similarity)
		}
	}

}
