package main

type TokenHandler struct{}

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{}
}

func (h *TokenHandler) generate() error {
	return nil
}

func (h *TokenHandler) validate() error {
	return nil
}

func (h *TokenHandler) delete() error {
	return nil
}
