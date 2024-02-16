package user

import (
	"cp23kk1/common/databases"
	userRepo "cp23kk1/modules/repository/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")

	userGroup.GET("/", getAllUsersHandler)
	userGroup.POST("/", createUserHandler)
	userGroup.GET("/:id", getUserHandler)
	userGroup.PUT("/:id", updateUserHandler)
	userGroup.DELETE("/:id", deleteUserHandler)
}

func getAllUsersHandler(c *gin.Context) {
	userRepository := userRepo.NewUserRepository(databases.GetDB())
	users, err := userRepository.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func createUserHandler(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRepository := userRepo.NewUserRepository(databases.GetDB())

	newUser := &databases.UserModel{
		Email:            userModelValidator.Email,
		Role:             userModelValidator.Role,
		DisplayName:      &userModelValidator.DisplayName,
		Image:            userModelValidator.Image,
		IsPrivateProfile: userModelValidator.IsPrivateProfile,
	}
	_, err := userRepository.CreateUser(*newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func getUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userRepository := userRepo.NewUserRepository(databases.GetDB())

	user, err := userRepository.FindUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func updateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRepository := userRepo.NewUserRepository(databases.GetDB())

	err = userRepository.UpdateUser(id, userModelValidator.Email, userModelValidator.Role, userModelValidator.DisplayName, userModelValidator.Image, userModelValidator.IsPrivateProfile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func deleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userRepository := userRepo.NewUserRepository(databases.GetDB())

	err = userRepository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
