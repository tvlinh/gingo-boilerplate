package presenter

import "tvlinh/gingo-boilerplate/internal/modules/helloworld/model"

type Hello struct {
	Message string
}

func (h *Hello) FromModel(m *model.Hello) *Hello {
	return &Hello{Message: m.Message}
}
