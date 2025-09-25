package controllers

import "net/http"

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Get(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (c *UserController) Create() string {
	return "Created a user"
}

func (c *UserController) Update() string {
	return "Updated a user"
}

func (c *UserController) Delete() string {
	return "Deleted a user"
}
