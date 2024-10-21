package model

type OutputImmersionAnkiFormat struct {
	Image              string `csv:"image"`
	ReadingText        string `csv:"readingText"`
	Audio              string `csv:"audio"`
	AnswerText         string `csv:"answerText"`
	AnswerTextFurigana string `csv:"answerTextFurigana"`
	SortOrder          int    `csv:"sortOrder"`
	OriginalText       string `csv:"originalText"`
}

type OutputWaniKaniAnkiFormat struct {
	Title               string `csv:"title"`
	Meaning             string `csv:"meaning"`
	AlternativeMeanings string `csv:"alternative_meanings"`
	Reading             string `csv:"readings"`
	Index               int    `csv:"index"`
}
