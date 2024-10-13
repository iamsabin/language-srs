package main

import (
	"os"

	"github.com/jszwec/csvutil"

	"language-srs/wanikani"
)

// // API Driven
// func main() {
// 	allSubjects := wanikani.GetUnknownSubjects()
// 	// panic(len(allSubjects))
// 	loadedAnki := alreadyLoadedAnki()
//
// 	for _, subject := range allSubjects {
// 		if existsAlready(subject, loadedAnki) {
// 			continue
// 		}
// 		anki, _ := getContextSentences(subject)
//
// 		if len(anki) == 0 {
// 			fmt.Println(subject.Text)
// 			continue
// 		}
//
// 		createAnkiOutput(anki, "contextsentences")
// 	}
// }

// Constant Driven
func main() {
	allSubjects := getRecentMistakes()

	for _, subject := range allSubjects {
		anki, _ := getContextSentences(subject)

		createAnkiOutput(anki, "recentmistakescontextsentencesv3")
	}
}

func existsAlready(subject wanikani.Subject, ankiLoaded []AnkiFormat) bool {
	for _, v := range ankiLoaded {
		if subject.Text == v.OriginalText {
			return true
		}
	}

	return false
}

func alreadyLoadedAnki() []AnkiFormat {
	data, err := os.ReadFile("wanikani/contextsentences.csv")
	if err != nil {
		panic(err)
	}

	var allAnki []AnkiFormat
	err = csvutil.Unmarshal(data, &allAnki)
	if err != nil {
		panic(err)
	}

	return allAnki
}
