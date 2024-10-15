package immersion

type Response struct {
	Data []struct {
		CategoryCount struct {
			Anime      int `json:"anime"`
			Drama      int `json:"drama"`
			Games      int `json:"games"`
			Literature int `json:"literature"`
			News       int `json:"news"`
		} `json:"category_count"`
		ExactMatch string `json:"exact_match"`
		Examples   []struct {
			Category             string `json:"category"`
			ImageUrl             string `json:"image_url"`
			Sentence             string `json:"sentence"`
			SentenceWithFurigana string `json:"sentence_with_furigana"`
			SoundUrl             string `json:"sound_url"`
			Translation          string `json:"translation"`
		} `json:"examples"`
	} `json:"data"`
}
