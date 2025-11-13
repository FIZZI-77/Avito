package handler

import (
	"avito/src/core/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", h.Ping)

	team := router.Group("/team")
	{
		team.POST("/add", h.CreateTeam)
		team.GET("/get", h.GetTeam)
	}

	users := router.Group("/users")
	{
		users.POST("/setIsActive", h.SetIsActive)
		users.GET("/getReview", h.GetReviewAssignments)
	}

	pr := router.Group("/pullRequest")
	{
		pr.POST("/create", h.CreatePullRequest)
		pr.POST("/merge", h.MergePullRequest)
		pr.POST("/reassign", h.ReassignReviewer)
	}

	return router
}
