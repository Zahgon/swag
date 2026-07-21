package controller

type Controller struct {
}

func NewController() *Controller { _ = "STUB: not implemented"; return nil }

type Message struct {
	Message string `json:"message" example:"message"`
}
