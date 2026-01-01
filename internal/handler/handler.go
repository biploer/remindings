package handler

import (
	"net/http"
)

type Dependences struct {
}

type router interface {
	http.Handler
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

func RegisterRoutes(router router, deps Dependences) {
	// router.HandleFunc("GET /", handle(main))
	// router.HandleFunc("GET /hello-my-love", handle(hello))
}
