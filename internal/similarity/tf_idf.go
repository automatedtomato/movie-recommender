package similarity

import (
	"math"
	"movie-recommender/internal/models"
)

// Term Frequency (TF) Concept:
// TF measures how frequently a word appears in a document
// It helps identify important words in a specific document
// Formula: TF(word) = (Number of times word appears) / (Total number of words in document)

// Here's a task for you:
//   - Create a function CalculateTermFrequency(words []string) map[string]float64
//   - The function should:
//   - Take a slice of preprocessed words
//   - Return a map where:
//     > Key is the word
//     > Value is its term frequency (count of word / total words)
func CalculateTermFrequency(words []string) map[string]float64 {
	// Count word occurrences
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}

	// Calculate term frequency
	totalWords := len(words)
	termFrequency := make(map[string]float64)

	for word, count := range wordCounts {
		termFrequency[word] = float64(count) / float64(totalWords)
	}

	return termFrequency
}

// DF Formula:
// IDF(word) = log(Total Number of Documents / Number of Documents Containing the Word)
// We'll need a new function that can:
//  - Take multiple movie descriptions
//  - Calculate how many documents each word appears in
//  - Calculate the IDF for each word

// CalculateIDF calculates Inverse Document Frequency for all documents
func CalculateIDF(allDocuments [][]string) map[string]float64 {
	// Total number of documents in the corpus
	totalDocuments := len(allDocuments)

	// Map to track number of documents containing each word
	wordDocumentCount := make(map[string]int)

	// Iterate through each document to count word document frequency
	for _, doc := range allDocuments {
		// Create a set of unique words in current document
		uniqueWords := make(map[string]bool)
		for _, word := range doc {
			uniqueWords[word] = true
		}

		// Increment document count for each unique word
		for word := range uniqueWords {
			wordDocumentCount[word]++
		}
	}

	// Calculate IDF scores for each word
	idfScores := make(map[string]float64)
	for word, documentCount := range wordDocumentCount {
		// IDF formula: log(Total Documents / Documents with word)
		idfScores[word] = math.Log(float64(totalDocuments) / float64(documentCount))
	}

	return idfScores
}

// CalculateTFIDF computes Term Frequency-Inverse Document Frequency for movie descriptions
func CalculateTFIDF(movies []*models.Movie) map[string]map[string]float64 {
	// Preprocess all movie descriptions
	preprocessedDocuments := make([][]string, len(movies))
	for i, movie := range movies {
		preprocessedDocuments[i] = PreProcess(movie.Description)
	}

	// Calculate IDF scores across all documents
	idfScores := CalculateIDF(preprocessedDocuments)

	// Container for TF-IDF vectors for each movie
	tfidfVectors := make(map[string]map[string]float64)

	// Compute TF-IDF for each movie description
	for i, document := range preprocessedDocuments {
		// Get corresponding movie title
		movieTitle := movies[i].Title

		// Calculate term frequencies for current document
		tfScores := CalculateTermFrequency(document)

		// Create TF-IDF vector for current document
		tfidfVector := make(map[string]float64)
		for word, tf := range tfScores {
			// TF-IDF is product of Term Frequency and Inverse Document Frequency
			tfidfVector[word] = tf * idfScores[word]
		}

		// Store TF-IDF vector for current movie
		tfidfVectors[movieTitle] = tfidfVector
	}

	return tfidfVectors
}

// CosineSimilarity calculates similarity between two TF-IDF vectors
func CosineSimilarity(vec1, vec2 map[string]float64) float64 {
	// Initialize dot product and magnitude calculations
	dotProduct := 0.0
	magnitude1 := 0.0
	magnitude2 := 0.0

	// Calculate dot product and first vector's magnitude
	for word, tfidf := range vec1 {
		// Compute dot product by multiplying corresponding TF-IDF values
		dotProduct += tfidf * vec2[word]
		// Compute squared magnitude of first vector
		magnitude1 += tfidf * tfidf
	}

	// Compute magnitude of second vector
	for _, tfidf := range vec2 {
		// Compute squared magnitude of second vector
		magnitude2 += tfidf * tfidf
	}

	// Take square root to get actual vector magnitudes
	magnitude1 = math.Sqrt(magnitude1)
	magnitude2 = math.Sqrt(magnitude2)

	// Handle zero magnitude vectors
	if magnitude1 == 0 || magnitude2 == 0 {
		return 0
	}

	// Compute and return cosine similarity
	return dotProduct / (magnitude1 * magnitude2)
}
