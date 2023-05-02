package handler

import (
	"gin/src/usecase"
	"sync"
)

var once = sync.Once{}

type Handler struct {
	uc *usecase.Usecase
}

func InitHandler(uc *usecase.Usecase) *Handler {
	r := Handler{}
	once.Do(func() {
		r.uc = uc
	})
	return &r
}
