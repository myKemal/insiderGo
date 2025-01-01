package handler

import (
	"encoding/json"
	"github.com/myKemal/insiderGo/app/model"
	"net/http"
)

// GetList godoc
// @Summary Get keys with pagination (POST)
// @Description Retrieve data from TempMemory using skip and limit
// @Tags TempMemory
// @Accept json
// @Produce json
// @Param body body model.TempListPayload true "Pagination parameters"
// @Success 200 {array} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /temp/getList [post]
// @Param Authorization header string true "<token>"
func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {

	var reqBody model.TempListPayload
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if reqBody.Skip < 0 {
		reqBody.Skip = 0
	}
	if reqBody.Limit <= 0 {
		reqBody.Limit = 10
	}

	keys, err := h.Temp.List(reqBody.Skip, reqBody.Limit)
	if err != nil {
		http.Error(w, "Failed to retrieve keys", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(keys)
	if err != nil {
		return
	}
}

// GetAllList godoc
// @Summary Get keys with pagination (GET)
// @Description Retrieve data from TempMemory
// @Tags TempMemory
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /temp/getAllList [get]
// @Param Authorization header string true "<token>"
func (h *Handler) GetAllList(w http.ResponseWriter, r *http.Request) {

	keys, err := h.Temp.AllList()
	if err != nil {
		http.Error(w, "Failed to retrieve keys", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(keys)
	if err != nil {
		return
	}
}
