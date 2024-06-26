package auth

import (
	"cp23kk1/common"
	"cp23kk1/common/config"
	"cp23kk1/common/databases"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, userId uint) {
	var myUserModel databases.UserModel
	if userId != 0 {
		db := databases.GetDB()
		db.First(&myUserModel, userId)
	}
	c.Set("userId", userId)
	c.Set("userModel", myUserModel)
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//
//	r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool, tokenName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Log or handle the Origin information as needed
		fmt.Printf("Request from Origin: %s\n", origin)
		UpdateContextUserModel(c, 0)
		config, _ := config.LoadConfig()
		token, err := c.Cookie(tokenName)
		key := ""
		if tokenName == "access_token" {
			key = config.AccessTokenPrivateKey
		} else {
			key = config.RefreshTokenPrivateKey
		}
		subject, err := common.ValidateToken(token, key)
		if err != nil {
			if auto401 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, common.ConvertVocaVerseResponse(common.VocaVerseStatusResponse{Status: "failed", Message: err.Error()}, map[string]interface{}{}))
			}
			return
		}

		userId := subject
		UpdateContextUserModel(c, userId)

	}
}
