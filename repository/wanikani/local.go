package wanikani

import (
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

func newLocalRepository() localRepository {
	path := "repository/wanikani/data/known-word.csv"
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var allSubjects []Subject
	err = csvutil.Unmarshal(data, &allSubjects)
	if err != nil {
		panic(err)
	}
	return local{
		path:        path,
		allSubjects: allSubjects,
	}
}

type local struct {
	path        string
	allSubjects []Subject
}

func (r local) getObjectValue(objectID int) (*WaniKaniSubject, error) {
	for _, v := range r.allSubjects {
		if v.ID == objectID {
			return &WaniKaniSubject{
				ID: v.ID,
				Data: struct {
					Characters string `json:"characters"`
					Meanings   []struct {
						Meaning string `json:"meaning"`
					} `json:"meanings"`
					Readings []struct {
						Reading string `json:"reading"`
					} `json:"readings"`
				}{Characters: v.Text},
			}, nil
		}
	}

	return nil, fmt.Errorf("object %d not found", objectID)
}

func (r local) setObjectValue(objectID int, objectValue string) error {
	r.allSubjects = append(r.allSubjects, Subject{
		ID:   objectID,
		Text: objectValue,
	})
	val, err := csvutil.Marshal(r.allSubjects)
	if err != nil {
		panic(err)
	}

	// TODO: append to a file
	file, err := os.Create(r.path)
	if err != nil {
		fmt.Printf("Failed to open or create file: %v\n", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(val)
	if err != nil {
		panic(err)
	}

	return nil
}
