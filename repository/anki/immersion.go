package anki

import (
	"fmt"
	"os"

	"github.com/jszwec/csvutil"

	"language-srs/model"
	"language-srs/repository"
)

type ankiRepository struct {
}

func (a ankiRepository) CreateImmersionDecks(output []model.ImmersionAnkiFormat,
	filename string) {
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

func NewAnkiRepository() repository.AnkiRepository {
	return ankiRepository{}
}
