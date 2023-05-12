package helper

import (
	"encoding/json"
	"net/http"
)

func RespError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(struct{ Message string }{msg})
}
