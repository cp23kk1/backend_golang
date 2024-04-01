package users

import (
	"fmt"
	"net/http"
	"strconv"

	"cp23kk1/common"
	"cp23kk1/common/databases"
	"cp23kk1/modules/auth"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.Use(auth.AuthMiddleware(true, "access_token"))
	users.GET("/profile", GetProfile)
	users.GET("/profile/:id", GetProfile)
	users.GET("/statistic", GetStatistic)
}

func GetProfile(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	var user *databases.UserModel = nil
	if err != nil {
		user, err = getUser(c.MustGet("userId").(uint))
	} else {

		user, err = getUser(uint(userId))
	}
	if err != nil {
		c.JSON(http.StatusNotFound, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "failed", Message: err.Error()}, map[string]interface{}{}))
		return
	}
	fmt.Println("userprofile ", c.MustGet("userId").(uint))
	serealizer := UserSerializer{c, *user}
	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "success", Message: "Get profile user successfully."}, serealizer.Response()))

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

func GetStatistic(c *gin.Context) {
	// userId, err := strconv.Atoi(c.Param("id"))
	userId, _ := c.MustGet("userId").(uint)

	getStatisticService(c, int(userId))

}

func UpdateProfile(c *gin.Context) {
	userId, _ := c.MustGet("userId").(uint)
	// userId := uint(5)
	changeDisplayNameValidator := NewChangeDisplayNameValidator()
	if err := changeDisplayNameValidator.Bind(c); err != nil {

		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: err.Error(), Status: "failed"}, map[string]interface{}{}))
		return
	}
	if err := updateUser(userId, changeDisplayNameValidator); err != nil {

		c.JSON(http.StatusBadRequest, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: err.Error(), Status: "failed"}, map[string]interface{}{}))
		return
	}

	c.JSON(http.StatusOK, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Message: "Update user success", Status: "successfully"}, map[string]interface{}{}))

}
