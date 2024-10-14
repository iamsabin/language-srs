package wanikani

type Subject struct {
	ID   int    `csv:"id"`
	Text string `csv:"text"`
}

type WaniKaniSubject struct {
	ID     int    `json:"id"`
	Object string `json:"object"`
	Data   struct {
		Characters string `json:"characters"`
		Meanings   []struct {
			Meaning string `json:"meaning"`
		} `json:"meanings"`
		Readings []struct {
			Reading string `json:"reading"`
		} `json:"readings"`
	} `json:"data"`
}
