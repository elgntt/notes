package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/elgntt/notes/internal/api/middleware"
	"github.com/elgntt/notes/internal/api/validation"
)

type Handler struct {
	taskService taskService
	logger      logger
}

var (
	errInvalidTaskIdErr = errors.New("невалидный параметр taskId")
	errTaskIdNotSet     = errors.New("не указан taskId")
)

func New(taskService taskService, logger logger) (*gin.Engine, error) {
	h := Handler{
		taskService: taskService,
		logger:      logger,
	}

	if err := validation.Set(); err != nil {
		return nil, err
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	r.Use(middleware.RegisterJSONTag())
	r.Use(middleware.ValidationErrors())

	r.POST("/api/tasks", h.CreateTask)
	r.PUT("/api/tasks/:taskId", h.UpdateTask)
	r.DELETE("/api/tasks/:taskId", h.DeleteTask)
	r.GET("/api/tasks", h.GetAllTasks)
	r.GET("/api/tasks/:taskId", h.GetTask)

	return r, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}
}

func parseTaskID(taskIdParam string) (int, error) {
	if taskIdParam == "" {
		return 0, errTaskIdNotSet
	}

	noteId, err := strconv.Atoi(taskIdParam)
	if err != nil {
		return 0, errInvalidTaskIdErr
	}

	return noteId, nil
}
