package controllers

import (
	"backend-crowd-funding/campaign"
	"backend-crowd-funding/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignControllers struct {
	service campaign.Service
}

func NewCampaignControllers(service campaign.Service) *campaignControllers {
	return &campaignControllers{service}
}

func (h *campaignControllers) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("Errpr get campaigns", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	response := helper.APIResponse("Success get campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
