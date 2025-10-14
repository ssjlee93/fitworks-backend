package controllers

import "net/http"

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
