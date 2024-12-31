package handler

import (
	"encoding/json"
	"github.com/myKemal/mongoApi/app/dtos"
	"net/http"
)

// FetchUnsentMessages godoc
// @Summary Fetch unsent messages
// @Description Get all messages with a sending status of "not_sent"
// @Tags Messages
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dtos.MessageDTO
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /unsent-messages [get]
// @Param Authorization header string true "<token>"
func (h *Handler) FetchUnsentMessages(w http.ResponseWriter, r *http.Request) {
	daos, err := h.Mongo.FetchUnsentMessages()
	if err != nil {
		http.Error(w, "Failed to fetch unsent messages", http.StatusInternalServerError)
		return
	}

	var myDtos []dtos.MessageDTO
	for _, dao := range daos {
		myDtos = append(myDtos, dtos.MessageDTO{
			MessageContent: dao.MessageContent,
			RecipientPhone: dao.RecipientPhone,
			SendingStatus:  dao.SendingStatus,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(myDtos)
	if err != nil {
		return
	}
}
