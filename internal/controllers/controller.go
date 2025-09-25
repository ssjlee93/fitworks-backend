package controllers

type Controller interface {
	Get() any
	Create() any
	Update() any
	Delete() any
}
