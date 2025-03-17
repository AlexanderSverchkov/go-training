package httphandlers

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

type RandomIntHandler struct{}

func NewRandomIntHandler(router *http.ServeMux) {
	handler := &RandomIntHandler{}
	router.HandleFunc("/random", handler.GetRandomInt())
}

func (handler *RandomIntHandler) GetRandomInt() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		randInt := rand.IntN(10)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(randInt)))
	}
}
