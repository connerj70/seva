package sandwich

// Sandwich holds all the data needed to construct a perfect sandwich
type Sandwich struct {
	Meat       string   `json:",omitempty"`
	Vegetables []string `json:",omitempty"`
	Bread      string   `json:",omitempty"`
	Delicious  *bool    `json:",omitempty"`
}
