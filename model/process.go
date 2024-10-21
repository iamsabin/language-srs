package model

type Transliterate struct {
	Kanji    string
	Kana     string
	Meanings []string
}

type WaniKaniSubject struct {
	ID   int    `csv:"id"`
	Text string `csv:"text"`
}
