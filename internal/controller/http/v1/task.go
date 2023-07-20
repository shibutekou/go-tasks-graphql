package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vgekko/go-tasks-graphql/internal/entity"
	"github.com/vgekko/go-tasks-graphql/internal/usecase"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

type taskRoutes struct {
	uc  usecase.Task
	log *slog.Logger
}

func newTaskRoutes(handler *gin.RouterGroup, uc usecase.Task, log *slog.Logger) {
	r := &taskRoutes{
		uc:  uc,
		log: log,
	}

	task := handler.Group("/task")
	{
		task.POST("/", r.createTask)
		task.GET("/:id", r.getTaskByID)
		task.PUT("/:id", r.updateTask)
		task.DELETE("/:id", r.deleteTask)

		task.GET("/completed", r.getCompletedTasks)
		task.GET("/opened", r.getOpenedTasks)

		task.PUT("/complete/:id", r.completeTask)
		task.PUT("/reopen/:id", r.reopenTask)
	}
}

func (r *taskRoutes) createTask(c *gin.Context) {
	var input entity.TaskInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	taskID, err := r.uc.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": taskID,
	})
}

func (r *taskRoutes) getTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := r.uc.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (r *taskRoutes) getAllTasks(c *gin.Context) {
	tasks, err := r.uc.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string][]entity.Task{
		"tasks": tasks,
	})
}

func (r *taskRoutes) getCompletedTasks(c *gin.Context) {
	completedTasks, err := r.uc.GetCompleted()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string][]entity.Task{
		"tasks": completedTasks,
	})
}

func (r *taskRoutes) getOpenedTasks(c *gin.Context) {
	openedTasks, err := r.uc.GetOpened()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string][]entity.Task{
		"tasks": openedTasks,
	})
}

func (r *taskRoutes) completeTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = r.uc.Complete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (r *taskRoutes) reopenTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = r.uc.Reopen(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (r *taskRoutes) updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input entity.TaskInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := r.uc.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (r *taskRoutes) deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.uc.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
