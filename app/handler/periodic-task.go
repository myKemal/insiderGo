package handler

import (
	"encoding/json"
	"net/http"

	"github.com/myKemal/insiderGo/app/services"
)

type PeriodicTaskHandler struct {
	Service *services.PeriodicTaskService
}

func NewPeriodicTaskHandler(service *services.PeriodicTaskService) *PeriodicTaskHandler {
	return &PeriodicTaskHandler{Service: service}
}

// StartPeriodicTaskService godoc
// @Summary Start the periodic task service
// @Description Starts the periodic task service to process tasks at regular intervals
// @Tags PeriodicTask
// @Produce json
// @Success 200 {object} map[string]string
// @Router /task/start [post]
// @Param Authorization header string true "<token>"
func (h *PeriodicTaskHandler) StartPeriodicTaskService(w http.ResponseWriter, r *http.Request) {
	started, err := h.Service.Start()
	if err != nil {
		http.Error(w, "Failed to start the service", http.StatusInternalServerError)
		return
	}

	if started {
		json.NewEncoder(w).Encode(map[string]string{"message": "PeriodicTask service started successfully"})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "PeriodicTask service is already running"})
	}
}

// StopPeriodicTaskService godoc
// @Summary Stop the periodic task service
// @Description Stops the periodic task service and halts task processing
// @Tags PeriodicTask
// @Produce json
// @Success 200 {object} map[string]string
// @Router /task/stop [post]
// @Param Authorization header string true "<token>"
func (h *PeriodicTaskHandler) StopPeriodicTaskService(w http.ResponseWriter, r *http.Request) {
	stopped, err := h.Service.Stop()
	if err != nil {
		http.Error(w, "Failed to stop the service", http.StatusInternalServerError)
		return
	}

	if stopped {
		json.NewEncoder(w).Encode(map[string]string{"message": "PeriodicTask service stopped successfully"})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "PeriodicTask service is not running"})
	}
}
