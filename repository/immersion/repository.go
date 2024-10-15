package immersion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"language-srs/model"
)

func GetImmersionInfo(keyword model.WaniKaniSubject) ([]model.AnkiFormat,
	error) {
	// Define the API endpoint with the query parameters
	apiURL := fmt.Sprintf("https://api.immersionkit.com/look_up_dictionary?keyword=%s&sort=shortness&wk=18",
		keyword.Text)

	// Make the HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return nil, err
	}

	// Parse the JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Failed to parse JSON response: %v\n", err)
		return nil, err
	}

	var ankiFormats []model.AnkiFormat
	// Print the parsed data
	for _, item := range response.Data {

		for i, v := range item.Examples {
			if i > 2 {
				break
			}
			ankiFormats = append(ankiFormats, model.AnkiFormat{
				Image:              v.ImageUrl,
				ReadingText:        v.Sentence,
				Audio:              v.SoundUrl,
				AnswerTextFurigana: v.SentenceWithFurigana,
				AnswerText:         v.Translation,
				SortOrder:          keyword.ID,
				OriginalText:       keyword.Text,
			})
		}
	}

	return ankiFormats, nil
}
