package tangochou

// type TangoChou struct {
// 	Id       string    `json:"id"`
// 	Name     string    `json:"name"`
// 	Sessions []string  `json:"sessions"`
// 	Srs      SRS       `json:"srs"`
// 	Subjects []Subject `json:"subjects"`
// }
//
// type SRS struct {
// 	Configuration Configuration `json:"configuration"`
// 	Enabled       bool          `json:"enabled"`
// 	Items         []string      `json:"items"`
// }
//
// type Configuration struct {
// 	Interval     int    `json:"interval"`
// 	IntervalUnit string `json:"interval_unit"`
// }
//
// type Subject struct {
// 	Audios     []string   `json:"audios"`
// 	Characters Characters `json:"characters"`
// 	Hint       struct {
// 		Meaning string `json:"meaning"`
// 		Reading string `json:"reading"`
// 	} `json:"hint"`
// 	Id       string     `json:"id"`
// 	Meanings []Meanings `json:"meanings"`
// 	Mnemonic struct {
// 		Meaning string `json:"meaning"`
// 		Reading string `json:"reading"`
// 	} `json:"mnemonic"`
// 	Readings  []Readings `json:"readings"`
// 	Sentences []string   `json:"sentences"`
// 	Speech    struct {
// 		Parts []string `json:"parts"`
// 	} `json:"speech"`
// 	Slug   string `json:"slug"`
// 	Source string `json:"source"`
// 	Type   string `json:"type"`
// }
//
// type Characters struct {
// 	Images []string `json:"images"`
// 	Text   string   `json:"text"`
// }
//
// type Meanings struct {
// 	Accepted bool   `json:"accepted"`
// 	Primary  bool   `json:"primary"`
// 	Value    string `json:"value"`
// }
//
// type Readings struct {
// 	Accepted bool   `json:"accepted"`
// 	Primary  bool   `json:"primary"`
// 	Type     string `json:"type"`
// 	Value    string `json:"value"`
// }
