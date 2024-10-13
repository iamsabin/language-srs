package transliterate

import "strings"

type Gloss struct {
	// Pos   string `json:"pos"`
	Gloss string `json:"gloss"`
	// Info  string `json:"info,omitempty"`
}

func (g *Gloss) GetMeanings() []string {
	return strings.Split(g.Gloss, "; ")
}

type WordInfo struct {
	// Reading string        `json:"reading"`
	Text string `json:"text"`
	Kana string `json:"kana"`
	// Score   int           `json:"score"`
	// Seq     int           `json:"seq"`
	Gloss      []Gloss       `json:"gloss"`
	Conj       []Conjugation `json:"conj"`
	Compound   []string      `json:"compound,omitempty"`
	Components []WordInfo    `json:"components,omitempty"`
	Suffix     string        `json:"suffix,omitempty"`
}

type Conjugation struct {
	Prop    []Prop  `json:"prop"`
	Reading string  `json:"reading"`
	Gloss   []Gloss `json:"gloss"`
	ReadOk  bool    `json:"readok"`
}

type Prop struct {
	Pos  string `json:"pos"`
	Type string `json:"type"`
}
