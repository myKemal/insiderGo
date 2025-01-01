package handler

import (
	"encoding/json"
	"github.com/myKemal/insiderGo/app/dtos"
	"github.com/myKemal/insiderGo/app/services"
	"net/http"
)

// GenerateToken godoc
// @Summary Generate JWT Token
// @Description Generates a JWT token for testing or authentication
// @Tags Auth
// @Produce json
// @Success 200 {object} dtos.TokenResponse
// @Failure 500 {object} map[string]string
// @Router /generate-token [post]
func (h *Handler) GenerateToken(w http.ResponseWriter, r *http.Request) {

	tokenString, err := services.GenerateToken()
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dtos.TokenResponse{Token: tokenString})
	if err != nil {
		return
	}
}
