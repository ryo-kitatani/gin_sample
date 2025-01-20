package handler

import (
	"gin-sample/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoHandler(todoUseCase usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

type createTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

type updateRequest struct {
	Status string `json:"status" binding:"required"`
}

// Create godoc
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todo
// @Accept json
// @Produce json
// @Param request body createTodoRequest true "Create Todo Request"
// @Success 201 {object} Todo
// @Router /todos [post]
func (h *TodoHandler) Create(c *gin.Context) {
	var req createTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.todoUseCase.Create(req.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAll godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func (h *TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.todoUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetByID godoc
// @Summary Get todo by ID
// @Description Get todo by ID
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} Todo
// @Router /todos/{id} [get]
func (h *TodoHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := h.todoUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateStatus godoc
// @Summary Update todo status
// @Description Update todo status
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param request body updateRequest true "Update Request"
// @Success 200 {object} Todo
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.todoUseCase.UpdateStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// Delete godoc
// @Summary Delete todo
// @Description Delete todo
// @Tags todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 204
// @Router /todos/{id} [delete]
func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.todoUseCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
