package usecase

import "tvlinh/gingo-boilerplate/internal/modules/helloworld/model"

type HelloAdvance interface {
	Hello() *model.Hello
}

func NewHelloAdvance(name string) HelloAdvance {
	return &helloAdvance{Name: name}
}

type helloAdvance struct {
	Name string
}

func (h *helloAdvance) Hello() *model.Hello {
	return &model.Hello{Message: "Hello World, " + h.Name}
}
