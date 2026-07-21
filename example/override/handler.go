package main

import (
	"database/sql"
	"net/http"
)

type ProductUpdates struct {
	Type        sql.NullString `json:"type"`
	Description sql.NullString `json:"description"`
	Stock       sql.NullInt64  `json:"stock"`
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
