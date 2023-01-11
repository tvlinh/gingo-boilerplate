package usecase

import "tvlinh/gingo-boilerplate/internal/modules/helloworld/model"

type HelloAdvance interface {
	Hello(name string) *model.Hello
}

func NewHelloAdvance() HelloAdvance {
	return &helloAdvance{}
}

type helloAdvance struct {
}

func (h *helloAdvance) Hello(name string) *model.Hello {
	return &model.Hello{Message: "Hello World, " + name}
}
