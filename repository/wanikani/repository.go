package wanikani

import (
	"strconv"
	"strings"

	"language-srs/repository"
)

func NewRepository() repository.Repository {
	return repo{}
}

type remoteRepository interface {
	getObjectIDList(srsStage string, onlyVocab bool) ([]Assignment, error)
	getObjectValue(objectID int) (*WaniKaniSubject, error)
}

type localRepository interface {
	getObjectValue(objectID int) (*WaniKaniSubject, error)
	setObjectValue(objectID int, objectValue string) error
}

type repo struct {
	remoteRepository
	localRepository
}

func (r repo) GetKnownWords() ([]string, error) {
	knownAssignments := []string{
		strconv.Itoa(srsBurned), strconv.Itoa(srsEnlightened),
	}
	knownWordList, _ := r.remoteRepository.getObjectIDList(strings.Join(knownAssignments,
		","), false)

	var knownWords []string

	for _, v := range knownWordList {
		object, _ := r.localRepository.getObjectValue(v.Data.SubjectID)
		if object == nil {
			object, _ = r.remoteRepository.getObjectValue(v.Data.SubjectID)

			if object != nil {
				_ = r.localRepository.setObjectValue(v.Data.SubjectID,
					object.Data.Characters)
			}
		}

		if object != nil {
			knownWords = append(knownWords, object.Data.Characters)
		}
	}

	return knownWords, nil
}
