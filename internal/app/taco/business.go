package taco

type ServiceAdapter interface{}

type Business struct{ Service ServiceAdapter }
