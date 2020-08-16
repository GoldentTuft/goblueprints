package thesaurus

// Thesaurus は
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
