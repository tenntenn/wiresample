package handler

import (
	"io"
	"net/http"

	"github.com/morikuni/failure"
)

const (
	badRequest failure.StringCode = "BadRequest"
)

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(errToStatus(err))
	io.WriteString(w, errToMessage(err))
}

func errToStatus(err error) int {
	c, ok := failure.CodeOf(err)
	if !ok {
		return http.StatusInternalServerError
	}
	switch c {
	case badRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func errToMessage(err error) string {
	msg, ok := failure.MessageOf(err)
	if ok {
		return msg
	}
	return "Error:" + http.StatusText(errToStatus(err))
}
