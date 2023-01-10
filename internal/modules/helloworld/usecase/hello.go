package usecase

import "tvlinh/gingo-boilerplate/internal/modules/helloworld/model"

type Hello interface {
	Hello() *model.Hello
}

func NewHello() Hello {
	return &hello{}
}

type hello struct {
}

func (h *hello) Hello() *model.Hello {
	return &model.Hello{Message: "Hello World!"}
}
