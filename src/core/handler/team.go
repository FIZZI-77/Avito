package handler

import (
	"avito/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "BAD_REQUEST", Message: err.Error()},
		})
		return
	}

	createdTeam, err := h.svc.Team.CreateTeam(c.Request.Context(), &team)
	if err != nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "TEAM_EXISTS", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, models.TeamResponse{Team: *createdTeam})
}

func (h *Handler) GetTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "BAD_REQUEST", Message: "team_name query is required"},
		})
		return
	}

	team, err := h.svc.Team.GetTeam(c.Request.Context(), teamName)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "NOT_FOUND", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, models.TeamResponse{Team: *team})
}
