package sandwich

// Sandwich holds all the data needed to construct a perfect sandwich
type Sandwich struct {
	Meat       string   `json:",omitempty"`
	Vegetables []string `json:",omitempty"`
	Bread      string   `json:",omitempty"`
	Delicious  *bool    `json:",omitempty"`
}

func (s *Sandwich) isDelicious() bool {
	if s.Delicious == nil {
		return false
	}
	if *s.Delicious {
		return true
	}
	return false
}

