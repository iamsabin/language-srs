package main

import (
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

type AnkiFormat struct {
	Image              string `csv:"image"`
	ReadingText        string `csv:"readingText"`
	Audio              string `csv:"audio"`
	AnswerText         string `csv:"answerText"`
	AnswerTextFurigana string `csv:"answerTextFurigana"`
	SortOrder          string `csv:"sortOrder"`
	OriginalText       string `csv:"originalText"`
}

func createAnkiOutput(output []AnkiFormat, filename string) {
	for i := range output {
		output[i].Image = fmt.Sprintf("<img src=\"%s\">", output[i].Image)
		output[i].Audio = fmt.Sprintf("[sound:%s]", output[i].Audio)
	}

	val, err := csvutil.Marshal(output)
	if err != nil {
		panic(err)
	}

	filePath := "wanikani/" + filename + ".csv"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)
	if err != nil {
		fmt.Printf("Failed to open or create file: %v\n", err)
		return
	}

	defer file.Close()

	_, err = file.Write(val)
	if err != nil {
		panic(err)
	}
}
