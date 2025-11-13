package handler

import (
	"avito/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreatePullRequest(c *gin.Context) {
	var req models.CreatePullRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "BAD_REQUEST", Message: err.Error()},
		})
		return
	}

	pr, err := h.svc.PullRequest.CreatePullRequest(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "PR_EXISTS", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, models.PullRequestResponse{PR: *pr})
}

func (h *Handler) MergePullRequest(c *gin.Context) {
	var req models.MergePullRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "BAD_REQUEST", Message: err.Error()},
		})
		return
	}

	pr, err := h.svc.PullRequest.MergePullRequest(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "NOT_FOUND", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, models.PullRequestResponse{PR: *pr})
}

func (h *Handler) ReassignReviewer(c *gin.Context) {
	var req models.ReassignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "BAD_REQUEST", Message: err.Error()},
		})
		return
	}

	res, err := h.svc.PullRequest.ReassignReviewer(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: models.ErrorDetail{Code: "REASSIGN_ERROR", Message: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
