package controllers

type Engine struct {
}

func New() (*Engine, error) {
	return &Engine{}, nil
}
