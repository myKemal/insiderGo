package handler

import (
	"encoding/json"
	"github.com/myKemal/insiderGo/app/services"
	"net/http"
)

// GetUnsentMessagesHandler godoc
// @Summary Get unsent messages
// @Description Retrieve a list of unsent messages from the system
// @Tags Instance
// @Produce json
// @Success 200 {array} dtos.MessageDTO
// @Failure 500 {object} map[string]string
// @Router /instance/unsent-messages [get]
// @Param Authorization header string true "<token>"
func GetUnsentMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messageService := services.GetMessageService()
	messages := messageService.GetMessages()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
