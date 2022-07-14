package handlers

import (
	"awesomeProject1/database/helper"
	"awesomeProject1/models"
	"encoding/json"
	"log"
	"net/http"
)

func DeleteRow(writer http.ResponseWriter, request *http.Request) {
	var req models.DeleteUser
	decoder := json.NewDecoder(request.Body)
	addErr := decoder.Decode(&req)

	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := helper.DeleteUser(req.ID)
	if err == nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf(req.ID)
}
