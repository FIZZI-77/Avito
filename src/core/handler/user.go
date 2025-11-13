package handler

import (
	"avito/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SetIsActive(c *gin.Context) {
	var req models.SetIsActiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "BAD_REQUEST", Message: err.Error()},
		})
		return
	}

	user, err := h.svc.User.SetIsActive(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "NOT_FOUND", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{User: *user})
}

func (h *Handler) GetReviewAssignments(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "BAD_REQUEST", Message: "user_id query is required"},
		})
		return
	}

	reviews, err := h.svc.User.GetReviewAssignments(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "NOT_FOUND", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
