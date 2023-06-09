package helper

import (
	"encoding/json"
	"net/http"
)

func RespError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct{ Message string }{msg})
}

type Response struct {
	Status int
	Err    error
}

func (r Response) Error() string {
	return r.Err.Error()
}
