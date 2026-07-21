package api

import (
	"net/http"
)

func GetStringByInt(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func GetStructArrayByString(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func Upload(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func AnonymousField() { _ = "STUB: not implemented"; return }

type Pet3 struct {
	ID int `json:"id"`
}
