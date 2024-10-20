package immersion

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"language-srs/model"
	"language-srs/repository"
)

const (
	MaxResults = 3

	WaniKaniLevel = 18

	CategoryDrama = "drama"
	CategoryAnime = "anime"
)

type repo struct {
	maxNumber     int
	category      string
	wanikaniLevel int
}

func NewRepository(waniKaniLevelLevel int) repository.ImmersionRepository {
	if waniKaniLevelLevel == 0 {
		waniKaniLevelLevel = WaniKaniLevel
	}
	return &repo{
		maxNumber:     MaxResults,
		category:      CategoryDrama,
		wanikaniLevel: waniKaniLevelLevel,
	}
}

func (r repo) GetImmersionInfo(keyword model.WaniKaniSubject) (
	[]model.ImmersionAnkiFormat, error) {
	apiURL := fmt.Sprintf(
		"https://api.immersionkit.com/look_up_dictionary?keyword=%s&category=%s&sort=shortness&wk=%d",
		keyword.Text, r.category, r.wanikaniLevel)

	// Make the HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		requestErr := fmt.Errorf(
			"failed to make request for url: %s, err: %w", apiURL, err)
		slog.Error(requestErr.Error())
		return nil, requestErr
	}
	defer func(Body io.ReadCloser) {
		bodyCloseErr := Body.Close()
		if bodyCloseErr != nil {
			bodyCloseError := fmt.Errorf(
				"failed to close body for url: %s,"+
					" err: %w", apiURL, bodyCloseErr)
			slog.Error(bodyCloseError.Error())
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		requestErr := fmt.Errorf(
			"failed to read response body for url: %s, err: %w", apiURL, err)
		slog.Error(requestErr.Error())
		return nil, requestErr
	}

	// Parse the JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		unmarshallErr := fmt.Errorf(
			"failed to parse JSON response for url: %s, "+
				"err: %w", apiURL, err)
		slog.Error(unmarshallErr.Error())
		return nil, unmarshallErr
	}

	var ankiFormats []model.ImmersionAnkiFormat
	// Print the parsed data
	for _, item := range response.Data {

		for i, v := range item.Examples {
			if i > r.maxNumber-1 {
				break
			}
			ankiFormats = append(
				ankiFormats, model.ImmersionAnkiFormat{
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

	// Check for anime category if drama is not found
	if len(ankiFormats) == 0 && r.category != CategoryAnime {
		r.category = CategoryAnime
		return r.GetImmersionInfo(keyword)
	}

	return ankiFormats, nil
}
