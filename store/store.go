package store

import "github.com/takumi616/go-english-vocabulary-api/entity"

var Vocabularies = &VocabularyStore{Vocabularies: map[entity.VocabularyID]*entity.Vocabulary{}}

type VocabularyStore struct {
	LastID       entity.VocabularyID
	Vocabularies map[entity.VocabularyID]*entity.Vocabulary
}

// Add a new vocabularry
func (vs *VocabularyStore) AddVocabulary(vocabulary *entity.Vocabulary) (entity.VocabularyID, error) {
	vs.LastID++
	vocabulary.ID = vs.LastID
	vs.Vocabularies[vocabulary.ID] = vocabulary
	return vocabulary.ID, nil
}
