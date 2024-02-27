package entity

import "time"

type VocabularyID int64

type Vocabulary struct {
	ID      VocabularyID `json:"vocabulary_id" db:"id"`
	Title   string       `json:"title" db:"title"`
	Example string       `json:"example" db:"example"`
	Created time.Time    `json:"created" db:"created"`
	Updated time.Time    `json:"updated" db:"updated"`
}
