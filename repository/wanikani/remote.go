package wanikani

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiBaseURL     = "https://api.wanikani.com/v2/subjects/"
	apiURL         = "https://api.wanikani.com/v2/assignments"
	apiKey         = ""
	srsBurned      = 9
	srsEnlightened = 8

	knownWordPath   = "wanikani/known-word.csv"
	unknownWordPath = "wanikani/unknown-word.csv"
)

func newRemoteRepository() remoteRepository {
	return remote{}
}

type remote struct {
}

func (r remote) getObjectIDList(srsStage string, onlyVocab bool) ([]Assignment,
	error) {
	var assignments []Assignment
	url := fmt.Sprintf("%s?srs_stages=%s", apiURL, srsStage)
	if onlyVocab {
		url += "&subject_types=kana_vocabulary,vocabulary"
	}
	for {
		// fmt.Printf("GET %s\n", url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+apiKey)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API request failed with status: %s",
				resp.Status)
		}

		var assignmentsResp AssignmentsResponse
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(body, &assignmentsResp); err != nil {
			return nil, err
		}

		assignments = append(assignments, assignmentsResp.Data...)
		if assignmentsResp.Pages.NextURL == "" {
			break
		}
		url = assignmentsResp.Pages.NextURL
	}
	return assignments, nil
}

func (r remote) getObjectValue(objectID int) (*WaniKaniSubject, error) {
	url := fmt.Sprintf("%s%d", apiBaseURL, objectID)
	fmt.Println("GET", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s",
			resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var subject WaniKaniSubject
	if err := json.Unmarshal(body, &subject); err != nil {
		return nil, err
	}

	time.Sleep(800 * time.Millisecond)

	return &subject, nil
}

type Assignment struct {
	Data struct {
		SrsStage  int `json:"srs_stage"`
		SubjectID int `json:"subject_id"`
	} `json:"data"`
}

type AssignmentsResponse struct {
	Data  []Assignment `json:"data"`
	Pages struct {
		NextURL string `json:"next_url"`
	} `json:"pages"`
}
