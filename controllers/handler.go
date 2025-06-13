package controller

import (
	"encoding/json"
	"fmt"
	"libr-simplified/db"
	"libr-simplified/model"
	"libr-simplified/moderator"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ts, _ := strconv.ParseInt(params["ts"], 10, 64)
	fmt.Printf("Get all messages with ts: %d\n", ts)
	json.NewEncoder(w).Encode(db.GetMessages(ts))
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := uuid.New()
	var requestBody struct {
		Content string `json:"content"`
	}
	_ = json.NewDecoder(r.Body).Decode(&requestBody)
	content := requestBody.Content

	var message model.Message
	message.Id = id.String()
	message.Content = content
	message.Timestamp = time.Now().Unix()
	message.Status = ""

	message.Status = moderator.SendMod(message)
	db.InsertMessage(message)

	response := model.MessageResponse{
		Id:        message.Id,
		Timestamp: message.Timestamp,
		Status:    message.Status,
	}
	fmt.Printf("Message added to database")
	json.NewEncoder(w).Encode(response)

}
