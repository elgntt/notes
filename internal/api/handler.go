package api

import (
	"context"

	"github.com/gin-gonic/gin"

	"task-manager/internal/api/categories"
	"task-manager/internal/api/projects"
	"task-manager/internal/api/tasks"
	"task-manager/internal/model/domain"
	"task-manager/internal/model/dto"
	validationpkg "task-manager/internal/pkg/validator"
)

type tasksService interface {
	Create(ctx context.Context, task dto.NewTask) (domain.Task, error)
	Update(ctx context.Context, updateTask dto.UpdateTask) error
	Delete(ctx context.Context, taskId int) error
	Get(ctx context.Context, categoryID int) ([]domain.Task, error)
	GetByID(ctx context.Context, taskId int) (domain.Task, error)
}

type categoriesService interface {
	Create(ctx context.Context, draft dto.NewCategory) (domain.Category, error)
	GetByID(ctx context.Context, id int) (domain.Category, error)
	FindByProjectID(ctx context.Context, projectId int) ([]domain.Category, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, draft dto.UpdateCategory) (domain.Category, error)
}

type projectsService interface {
	Create(ctx context.Context, project dto.Project) (domain.Project, error)
	GetByID(ctx context.Context, id int) (domain.Project, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, proj dto.Project, projectID int) (domain.Project, error)
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

func New(
	logger logger,
	tasksService tasksService,
	categoriesService categoriesService,
	projectsService projectsService,
) (*gin.Engine, error) {
	v, err := validationpkg.New()
	if err != nil {
		return nil, err
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	tasksAPI := tasks.NewAPI(logger, tasksService, v)
	api := r.Group("/api")
	api.POST("/tasks", tasksAPI.Create)
	api.PUT("/tasks/:taskId", tasksAPI.Update)
	api.DELETE("/tasks/:taskId", tasksAPI.Delete)
	api.GET("/categories/:categoryId/tasks", tasksAPI.List)
	api.GET("/tasks/:taskId", tasksAPI.Get)

	categoriesAPI := categories.NewAPI(logger, categoriesService, v)
	api.POST("/categories", categoriesAPI.Create)
	api.GET("/categories/:categoryId", categoriesAPI.Get)
	api.GET("/projects/:projectId/categories", categoriesAPI.List)
	api.DELETE("/categories/:categoryId", categoriesAPI.Delete)
	api.PUT("/categories/:categoryId", categoriesAPI.Update)

	projectsAPI := projects.NewAPI(logger, projectsService, v)
	api.POST("/projects", projectsAPI.Create)
	api.GET("/projects/:projectId", projectsAPI.Get)
	api.DELETE("/projects/:projectId", projectsAPI.Delete)
	api.PUT("/projects/:projectId", projectsAPI.Update)

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
