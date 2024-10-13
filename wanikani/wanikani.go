package wanikani

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jszwec/csvutil"
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

func getAssignments(srsStage string, onlyVocab bool) ([]Assignment, error) {
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

var allSubjects []Subject

func GetSubjects() []Subject {
	knownAssignments := []string{
		strconv.Itoa(srsBurned), strconv.Itoa(srsEnlightened),
	}
	burnedAssignments, err := getAssignments(strings.Join(knownAssignments,
		","), false)
	if err != nil {
		fmt.Printf("Error getting burned assignments: %v\n", err)
		os.Exit(1)
	}

	loadSubjects(knownWordPath)

	// for _, v := range burnedAssignments {
	// 	fmt.Println(v.Data.SubjectID)
	// }

	getSubjects(burnedAssignments, knownWordPath)

	createKnownItems(knownWordPath)

	return allSubjects
}

func GetUnknownSubjects() []Subject {
	apprenticeStages := []string{"1", "2", "3", "4"}
	// guruStages := []string{"5", "6"}
	// masterStages := []string{"7"}

	// unknownStages := append(apprenticeStages, guruStages...)
	// unknownStages = append(unknownStages, masterStages...)

	unknownAssignments, err := getAssignments(strings.Join(apprenticeStages,
		","), false)
	if err != nil {
		fmt.Printf("Error getting burned assignments: %v\n", err)
		os.Exit(1)
	}

	loadSubjects(unknownWordPath)

	getSubjects(unknownAssignments, unknownWordPath)

	createKnownItems(unknownWordPath)

	return allSubjects
}

func loadSubjects(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = csvutil.Unmarshal(data, &allSubjects)
	if err != nil {
		panic(err)
	}
}

func getSubjects(assignments []Assignment, path string) {
	for _, v := range assignments {
		if !hasLocally(strconv.Itoa(v.Data.SubjectID)) {
			res, err := getSubjectByID(v.Data.SubjectID)
			if err != nil {
				createKnownItems(path)
				panic(err)
			}

			if res == nil {
				continue
			}
			allSubjects = append(allSubjects, Subject{
				ID:   strconv.Itoa(res.ID),
				Text: res.Data.Characters,
			})
		}
	}
}

func hasLocally(id string) bool {
	for _, v := range allSubjects {
		if v.ID == id {
			return true
		}
	}

	return false
}

func getSubjectByID(subjectID int) (*WaniKaniSubject, error) {
	url := fmt.Sprintf("%s%d", apiBaseURL, subjectID)
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

func createKnownItems(path string) {
	val, err := csvutil.Marshal(allSubjects)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to open or create file: %v\n", err)
		return
	}

	defer file.Close()

	_, err = file.Write(val)
	if err != nil {
		panic(err)
	}
}
