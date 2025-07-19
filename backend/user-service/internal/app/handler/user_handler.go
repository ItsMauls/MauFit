package handler

import (
	"net/http"
	"user-service/internal/usecase"
	"user-service/pkg/util"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uc,
	}
}

type CreateUserInput struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	FingerprintID string `json:"fingerprint_id"`
}

func (h *UserHandler) VerifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
		return
	}
	// Validasi token, misal dengan JWT
	user, err := h.userUsecase.VerifyToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := util.APIResponse("Input not valid", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.userUsecase.CreateUser(input.Name, input.Email, input.FingerprintID)
	if err != nil {
		response := util.APIResponse(err.Error(), http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("User Created", http.StatusCreated, user)
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userUsecase.GetUserByID(id)
	if err != nil {
		response := util.APIResponse("User not found", http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := util.APIResponse("User found", http.StatusOK, user)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		response := util.APIResponse("Failed to get users", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("Users found", http.StatusOK, users)
	c.JSON(http.StatusOK, response)
}

type RegisterUserInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	PhotoProfileURL string `json:"photo_profile_url"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Bio             string `json:"bio"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := util.APIResponse("Input not valid", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.userUsecase.RegisterUser(input.Name, input.Email, input.Password, input.PhotoProfileURL, input.Address, input.Phone, input.Bio)
	if err != nil {
		response := util.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := util.APIResponse("User registered", http.StatusCreated, user)
	c.JSON(http.StatusCreated, response)
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := util.APIResponse("Input not valid", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.userUsecase.AdminLogin(input.Email, input.Password)
	if err != nil {
		response := util.APIResponse(err.Error(), http.StatusUnauthorized, nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	response := util.APIResponse("Login successful", http.StatusOK, gin.H{"token": token})
	c.JSON(http.StatusOK, response)
}

type UpdateUserProfileInput struct {
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Bio             string `json:"bio"`
	PhotoProfileURL string `json:"photo_profile_url"`
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	id := c.Param("id")
	var input UpdateUserProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response := util.APIResponse("Input not valid", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	profile, err := h.userUsecase.UpdateUserProfile(id, input.Address, input.Phone, input.Bio, input.PhotoProfileURL)
	if err != nil {
		response := util.APIResponse(err.Error(), http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := util.APIResponse("User profile updated", http.StatusOK, profile)
	c.JSON(http.StatusOK, response)
}
