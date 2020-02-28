package sandwich

type ServiceAdapter interface{}

type Business struct{ Service ServiceAdapter }
