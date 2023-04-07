package controller

type Context interface {
	AbortWithStatusJSON(code int, jsonObj interface{})
}
