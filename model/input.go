package model

type InputEnglishToJapanese struct {
	Words          []string
	OutputFilename string
}

type InputTransliterate struct {
	Japanese string `csv:"Japanese"`
	English  string `csv:"English"` // Not needed
}
