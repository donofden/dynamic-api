package controllers

import (
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

// GetConnection exported
var GetConnection = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	table := vars["table"]
	data := models.GetTableDataByID(table, id)
	resp := u.Message(true, "success")
	resp["table"] = table
	resp["data"] = data
	u.Respond(w, resp)
}
