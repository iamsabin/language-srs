package model

type Transliterate struct {
	Kanji    string
	Kana     string
	Meanings []string
}

type AnkiFormat struct {
	Image              string `csv:"image"`
	ReadingText        string `csv:"readingText"`
	Audio              string `csv:"audio"`
	AnswerText         string `csv:"answerText"`
	AnswerTextFurigana string `csv:"answerTextFurigana"`
	SortOrder          int    `csv:"sortOrder"`
	OriginalText       string `csv:"originalText"`
}

type WaniKaniSubject struct {
	ID   int    `csv:"id"`
	Text string `csv:"text"`
}
