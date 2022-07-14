package handlers

import (
	"awesomeProject1/database/helper"
	"awesomeProject1/models"
	"encoding/json"
	"log"
	"net/http"
)

func UpdateRow(writer http.ResponseWriter, request *http.Request) {
	var req models.UpdateUser
	decoder := json.NewDecoder(request.Body)
	addErr := decoder.Decode(&req)
	log.Printf(req.Name)
	if addErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, err := helper.UpdateUser(req.Name, req.Email, req.ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, userErr := helper.GetUser(userID)
	if userErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
}
