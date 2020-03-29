package sandwich

import "github.com/connerj70/seva/internal/cerr"

type ServiceAdapter interface{}

type Business struct{ Service ServiceAdapter }

func (b *Business) Post(sandwich *Sandwich) error {
	// validate
	if !sandwich.isDelicious() {
		return cerr.NewInternalError(404, "Sandwich Validation", "sandwich must have a delicious property of True", ErrNotDelicious)
	}
	return nil
}

func (b *Business) Put(*Sandwich) error {
	return nil
}
