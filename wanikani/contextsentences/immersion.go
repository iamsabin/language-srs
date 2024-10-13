package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"songs-srs/wanikani"
)

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

func getContextSentences(keyword wanikani.Subject) ([]AnkiFormat, error) {
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

	var ankiFormats []AnkiFormat
	// Print the parsed data
	for _, item := range response.Data {

		for i, v := range item.Examples {
			if i > 2 {
				break
			}
			ankiFormats = append(ankiFormats, AnkiFormat{
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
