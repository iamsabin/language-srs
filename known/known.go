package known

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jszwec/csvutil"
	"language-srs/anki"
	"language-srs/wanikani"
)

func GetSubjects() []wanikani.Subject {
	ankis := getInput()

	var allSubjects []wanikani.Subject

	for i, v := range ankis {
		allSubjects = append(allSubjects, wanikani.Subject{
			ID:   strconv.Itoa(i),
			Text: v.Title,
		})
		if v.Reading != "" {
			allSubjects = append(allSubjects, wanikani.Subject{
				ID:   strconv.Itoa(i),
				Text: v.Reading,
			})
		}
	}

	return allSubjects
}

func getInput() []anki.Anki {
	var input []anki.Anki

	// Directory containing the CSV files
	dir := "known/output"

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

			var singleInput []anki.Anki
			e := csvutil.Unmarshal(f, &singleInput)

			if e != nil {
				log.Fatalf("failed to read CSV file: %v", err)
			}

			for _, v := range singleInput {
				input = append(input, v)
			}
		}
	}

	return input
}
