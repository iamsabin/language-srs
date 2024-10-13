package wanikani

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Define the structure for the API response
type Assignment2 struct {
	ID     int    `json:"id"`
	Object string `json:"object"`
	Data   struct {
		SubjectID         int    `json:"subject_id"`
		Level             int    `json:"level"`
		IncorrectMeanings string `json:"incorrect_meanings"`
		IncorrectReadings string `json:"incorrect_readings"`
	} `json:"data"`
}

type AssignmentsResponse2 struct {
	Assignments []Assignment2 `json:"data"`
}

func GetRecentMistakes() {
	// Replace 'your_api_key_here' with your actual WaniKani API key
	url := "https://api.wanikani.com/v2/assignments"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add the API key to the request header
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse the response JSON
	var assignmentsResponse AssignmentsResponse2
	err = json.Unmarshal(body, &assignmentsResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print out the assignments with mistakes
	for _, assignment := range assignmentsResponse.Assignments {
		incorrectMeanings := strings.TrimSpace(assignment.Data.IncorrectMeanings)
		incorrectReadings := strings.TrimSpace(assignment.Data.IncorrectReadings)

		if incorrectMeanings != "" || incorrectReadings != "" {
			fmt.Printf("Assignment ID: %d\n", assignment.ID)
			fmt.Printf("Subject ID: %d\n", assignment.Data.SubjectID)
			fmt.Printf("Level: %d\n", assignment.Data.Level)
			fmt.Printf("Incorrect Meanings: %v\n", incorrectMeanings)
			fmt.Printf("Incorrect Readings: %v\n", incorrectReadings)
			fmt.Println()
		}
	}
}
