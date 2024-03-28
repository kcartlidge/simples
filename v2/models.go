package simples

// Entry ... A single entry within a section.
type Entry struct {
	Sequence      int
	Section       string
	Key, KeyUpper string
	Value         string
}

// Sections ... Complete section contents; unordered but with a sequence.
type Sections map[string][]Entry
