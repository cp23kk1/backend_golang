package users

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"cp23kk1/common"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/:id", UserRetrieve)
}

func UserRetrieve(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user, err := getUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("id", errors.New("Invalid ID")))
		return
	}
	var userResponse UserResponse
	err = mapstructure.Decode(user, &userResponse)
	if err != nil {
		fmt.Println("Error mapping data:", err)
	} 
	c.JSON(http.StatusOK, gin.H{"user": userResponse})

	// serializer := UserSerializer{c, user}
	// c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
	// users.UsersRegister(v1.Group("/users"))
	// v1.Use(users.AuthMiddleware(false))
	// articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	// articles.TagsAnonymousRegister(v1.Group("/tags"))

	// v1.Use(users.AuthMiddleware(true))
	// users.UserRegister(v1.Group("/user"))
	// users.ProfileRegister(v1.Group("/profiles"))
}