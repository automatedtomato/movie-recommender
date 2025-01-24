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

func CalculateIDF(allDocuments [][]string) map[string]float64 {
	// Total number of documents
	totalDocuments := len(allDocuments)

	// Count documents containing each word
	wordDocumentCount := make(map[string]int)

	// First pass: Count documents per word
	for _, doc := range allDocuments {
		// Use a set to count each word only once per document
		uniqueWordsInDoc := make(map[string]bool)

		for _, word := range doc {
			uniqueWordsInDoc[word] = true
		}

		// Increment count for words in this document
		for word := range uniqueWordsInDoc {
			wordDocumentCount[word]++
		}
	}

	// Calculate IDF for each word
	idfScores := make(map[string]float64)

	for word, documentCount := range wordDocumentCount {
		// IDF formula: log(Total Documents)/ Documents with word)
		idfScores[word] = math.Log(float64(totalDocuments) / float64(documentCount))
	}

	return idfScores
}

func CalculateTFIDF(movies []*models.Movie) map[string]map[string]float64 {
	// Step 1: Preprocess all descriptions
	preprocessedDocuments := make([][]string, len(movies))
	for i, movie := range movies {
		preprocessedDocuments[i] = PreProcess(movie.Description)
	}

	// Step 2: Calculate IDF for each word first
	idfScores := CalculateIDF(preprocessedDocuments)

	// Step 3: Calculate TF for each word in each document
	tfidfVectors := make(map[string]map[string]float64)

	for i, document := range preprocessedDocuments {
		// Get movie title as identifier
		movieTitle := movies[i].Title

		// Calculate TF for this document
		tfScores := CalculateTermFrequency(document)

		// Calculate TF-IDF for this document
		tfidfVector := make(map[string]float64)
		for word, tf := range tfScores {
			// TF-IDF = TF * IDF
			tfidfVector[word] = tf * idfScores[word]
		}

		tfidfVectors[movieTitle] = tfidfVector
	}

	return tfidfVectors
}

func CosineSimilarity(vec1, vec2 map[string]float64) float64 {
	// Calculate the dot product
	dotProduct := 0.0
	for word, tfidf := range vec1 {
		dotProduct += tfidf * vec2[word]
	}

	// Calculate the magnitude of both vectors
	magnitude1 := 0.0
	for _, val := range vec1 {
		magnitude1 = val * val
	}
	magnitude1 = math.Sqrt(magnitude1)

	magnitude2 := 0.0
	for _, val := range vec2 {
		magnitude2 = val * val
	}
	magnitude2 = math.Sqrt(magnitude2)

	// Handle cosine similarity
	return dotProduct / (magnitude1 * magnitude2)
}
