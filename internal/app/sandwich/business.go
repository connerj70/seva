package sandwich

type ServiceAdapter interface{}

type Business struct{ Service ServiceAdapter }

func (b *Business) Post(*Sandwich) error {
	return nil
}
