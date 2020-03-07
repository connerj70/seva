package taco

type BusinessAdapter interface{}

type Reciever struct{ Business BusinessAdapter }
