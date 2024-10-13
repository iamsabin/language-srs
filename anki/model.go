package anki

type Anki struct {
	Title               string `csv:"title"`
	Meaning             string `csv:"meaning"`
	AlternativeMeanings string `csv:"alternative_meanings"`
	Reading             string `csv:"readings"`
	Index               int    `csv:"index"`
}
