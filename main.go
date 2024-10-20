package main

import (
	"language-srs/model"
	"language-srs/service"
)

func main() {
	srv := service.NewService()

	inputEnToJP := model.InputEnglishToJapanese{
		Words:          []string{"すみません", "ありがとう"},
		OutputFilename: "recentmistakes-context-sentences-deck",
	}
	srv.CreateEnglishToJapaneseDeck(inputEnToJP)
}
