package passage

import (
	"cp23kk1/common"

	"github.com/gin-gonic/gin"
)

type PassageModelValidator struct {
	Title string `form:"title" json:"title" binding:"required,max=255"`
}

func NewPassageModelValidator() PassageModelValidator {
	return PassageModelValidator{}
}

func (s *PassageModelValidator) Bind(c *gin.Context) error {

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	return nil
}
