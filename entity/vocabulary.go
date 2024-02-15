package entity

import "time"

type VocabularyID int64

type Vocabulary struct {
	ID      VocabularyID `json:"vocabulary_id"`
	Title   string       `json:"title"`
	Example string       `json:"example"`
	Created time.Time    `json:"created"`
}
