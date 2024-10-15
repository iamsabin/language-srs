package tangochou

// import (
// 	"encoding/json"
// 	"os"
// 	"strconv"
// 	"strings"
//
// 	"language-srs/model"
//
// 	"github.com/google/uuid"
// )
//
// // OBSOLETE: No use any more
// func CreateSRSDeck(input []model.Transliterate, name string) {
//
// 	subjects := []Subject{}
//
// 	for _, v := range input {
// 		var meanings []Meanings
//
// 		for _, w := range v.Meanings {
// 			meanings = append(meanings, Meanings{
// 				Accepted: true,
// 				Primary:  true,
// 				Value:    w,
// 			})
// 		}
//
// 		sub := Subject{
// 			Audios: nil,
// 			Characters: Characters{
// 				Text: v.Kanji,
// 			},
// 			Id:       v.Kanji,
// 			Meanings: meanings,
// 			Slug:     v.Kanji,
// 			Source:   "ICHIMOE",
// 		}
//
// 		var readings []Readings
// 		for _, vK := range strings.Split(v.Kana, ",") {
// 			readings = append(readings, Readings{
// 				Accepted: true,
// 				Primary:  true,
// 				Type:     "UNKNOWN",
// 				Value:    vK,
// 			})
// 		}
// 		sub.Readings = readings
//
// 		subType := "VOCABULARY"
// 		if !isKanjiWord(v.Kanji) {
// 			subType = "RADICAL"
// 		}
// 		sub.Type = subType
//
// 		if !isAlreadyThere(subjects, sub) {
// 			subjects = append(subjects, sub)
// 		}
// 	}
//
// 	createTangoChouDeck(subjects, name)
// }
//
// func isAlreadyThere(subjects []Subject, subject Subject) bool {
// 	for _, s := range subjects {
// 		if s.Characters.Text == subject.Characters.Text {
// 			return true
// 		}
// 	}
//
// 	return false
// }
//
// func createTangoChouDeck(subjects []Subject, filename string) {
// 	splittedSubjects := splitTangoChouSubjects(subjects)
//
// 	for i, v := range splittedSubjects {
// 		if len(v) == 0 {
// 			continue
// 		}
// 		var decks []TangoChou
//
// 		name := filename + "-" + strconv.Itoa(i+1)
//
// 		deck := TangoChou{
// 			Id:   uuid.New().String(),
// 			Name: name,
// 			Srs: SRS{
// 				Configuration: Configuration{
// 					Interval:     1,
// 					IntervalUnit: "day",
// 				},
// 				Enabled: false,
// 				Items:   nil,
// 			},
// 		}
//
// 		deck.Subjects = v
//
// 		decks = append(decks, deck)
//
// 		val, err := json.Marshal(decks)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		valString := string(val)
//
// 		valString = strings.ReplaceAll(valString, "null", "[]")
//
// 		file, err := os.Create("output/" + name + ".json")
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		_, err = file.Write([]byte(valString))
// 		if err != nil {
// 			panic(err)
// 		}
// 		file.Close()
// 	}
//
// }
//
// func splitTangoChouSubjects(subjects []Subject) [][]Subject {
// 	maxContentInOneDeck := 10
//
// 	var splittedSubjects = make([][]Subject,
// 		len(subjects)/maxContentInOneDeck*2)
//
// 	var currentDeckCounter = 0
// 	for i, v := range subjects {
// 		if splittedSubjects[currentDeckCounter] == nil {
// 			splittedSubjects[currentDeckCounter] = []Subject{}
// 		}
//
// 		splittedSubjects[currentDeckCounter] = append(splittedSubjects[currentDeckCounter],
// 			v)
//
// 		if (i+1)%maxContentInOneDeck == 0 {
// 			currentDeckCounter++
// 		}
// 	}
//
// 	return splittedSubjects
// }
//
// // isKanji checks if a rune is a Kanji character
// func isKanji(r rune) bool {
// 	return (r >= 0x4E00 && r <= 0x9FBF) || // CJK Unified Ideographs
// 		(r >= 0x3400 && r <= 0x4DBF) || // CJK Unified Ideographs Extension A
// 		(r >= 0x20000 && r <= 0x2A6DF) || // CJK Unified Ideographs Extension B
// 		(r >= 0x2A700 && r <= 0x2B73F) || // CJK Unified Ideographs Extension C
// 		(r >= 0x2B740 && r <= 0x2B81F) || // CJK Unified Ideographs Extension D
// 		(r >= 0x2B820 && r <= 0x2CEAF) || // CJK Unified Ideographs Extension E
// 		(r >= 0xF900 && r <= 0xFAFF) || // CJK Compatibility Ideographs
// 		(r >= 0x2F800 && r <= 0x2FA1F) // CJK Compatibility Ideographs Supplement
// }
//
// // isKanjiWord checks if all characters in a word are Kanji
// func isKanjiWord(word string) bool {
// 	var hasKanji = false
//
// 	for _, r := range word {
// 		if isKanji(r) {
// 			hasKanji = true
// 		}
// 	}
// 	return hasKanji
// }
