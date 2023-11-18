package users

import (
	"net/http"
	"strconv"

	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/:id", UserRetrieve)
}

func UserRetrieve(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user, err := getUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: err.Error()}, map[string]interface{}{}))
		return
	}
	serealizer := UserSerializer{c, *user}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "success"}, map[string]interface{}{"users": serealizer.Response()}))

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
