package model

type Transliterate struct {
	Kanji    string
	Kana     string
	Meanings []string
}

type ImmersionAnkiFormat struct {
	Image              string `csv:"image"`
	ReadingText        string `csv:"readingText"`
	Audio              string `csv:"audio"`
	AnswerText         string `csv:"answerText"`
	AnswerTextFurigana string `csv:"answerTextFurigana"`
	SortOrder          int    `csv:"sortOrder"`
	OriginalText       string `csv:"originalText"`
}

type WaniKaniAnkiFormat struct {
	Title               string `csv:"title"`
	Meaning             string `csv:"meaning"`
	AlternativeMeanings string `csv:"alternative_meanings"`
	Reading             string `csv:"readings"`
	Index               int    `csv:"index"`
}

type WaniKaniSubject struct {
	ID   int    `csv:"id"`
	Text string `csv:"text"`
}

type InputEnglishToJapanese struct {
	Words          []string
	OutputFilename string
}
