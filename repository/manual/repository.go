package manual

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jszwec/csvutil"

	"language-srs/model"
	"language-srs/repository"
)

func NewRepository() repository.Repository {
	return repo{}
}

type repo struct {
}

func (r repo) GetKnownWords() ([]string, error) {
	return append(
		r.getKnownWordsFromJPToENStyle(),
		r.getKnownWordsFromWaniKaniStyle()...), nil
}

func (r repo) getKnownWordsFromWaniKaniStyle() []string {
	var input []model.OutputWaniKaniAnkiFormat

	// Directory containing the CSV files
	dir := "output/wanikanistyle"

	// Read the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	// Loop through all the files
	for _, file := range files {
		// Check if the file has a .csv extension
		if strings.HasSuffix(file.Name(), ".csv") {
			// Open the CSV file
			filePath := filepath.Join(dir, file.Name())
			f, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatalf("failed to open file: %v", err)
			}

			var singleInput []model.OutputWaniKaniAnkiFormat
			e := csvutil.Unmarshal(f, &singleInput)

			if e != nil {
				log.Fatalf("failed to read CSV file: %v", err)
			}

			for _, v := range singleInput {
				input = append(input, v)
			}
		}
	}

	var knownWords []string
	for _, v := range input {
		knownWords = append(knownWords, v.Title)
		if v.Reading != "" {
			knownWords = append(knownWords, v.Reading)
		}
	}

	return knownWords
}

func (r repo) getKnownWordsFromJPToENStyle() []string {
	var input []model.OutputImmersionAnkiFormat

	// Directory containing the CSV files
	dir := "output/jptoen"

	// Read the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	// Loop through all the files
	for _, file := range files {
		// Check if the file has a .csv extension
		if strings.HasSuffix(file.Name(), ".csv") {
			// Open the CSV file
			filePath := filepath.Join(dir, file.Name())
			f, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatalf("failed to open file: %v", err)
			}

			var singleInput []model.OutputImmersionAnkiFormat
			e := csvutil.Unmarshal(f, &singleInput)

			if e != nil {
				log.Fatalf("failed to read CSV file: %v", err)
			}

			for _, v := range singleInput {
				input = append(input, v)
			}
		}
	}

	var knownWords []string
	for _, v := range input {
		knownWords = append(knownWords, v.JPToLearn)
	}

	return knownWords
}
