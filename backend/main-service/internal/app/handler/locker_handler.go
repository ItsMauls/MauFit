package handler

import (
	"main-service/internal/usecase"
	"main-service/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LockerHandler struct {
	lockerUsecase usecase.LockerUsecase
}

func NewLockerHandler(uc usecase.LockerUsecase) *LockerHandler {
	return &LockerHandler{
		lockerUsecase: uc,
	}
}

type CreateLockerInput struct {
	LockerNumber uint    `json:"number" binding:"required"`
	IsUsed bool `json:"status" binding:"required"`
}

func (h *LockerHandler) CreateLocker(c *gin.Context) {
	var input CreateLockerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := util.APIResponse("Input not valid", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	locker, err := h.lockerUsecase.CreateLocker(input.LockerNumber, input.IsUsed)
	if err != nil {
		response := util.APIResponse(err.Error(), http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("Locker created", http.StatusCreated, locker)
	c.JSON(http.StatusCreated, response)
}

func (h *LockerHandler) GetLockerByID(c *gin.Context) {
	id := c.Param("id")
	locker, err := h.lockerUsecase.GetLockerByID(id)
	if err != nil {
		response := util.APIResponse("Locker not found", http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := util.APIResponse("Locker found", http.StatusOK, locker)
	c.JSON(http.StatusOK, response)
}

func (h *LockerHandler) GetAllLockers(c *gin.Context) {
	lockers, err := h.lockerUsecase.GetAllLockers()
	if err != nil {
		response := util.APIResponse("Failed to get lockers", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("Lockers found", http.StatusOK, lockers)
	c.JSON(http.StatusOK, response)
}
