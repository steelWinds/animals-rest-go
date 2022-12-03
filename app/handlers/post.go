package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostH[T any](orm *gorm.DB) func (*gin.Context) {
	handler := func (c *gin.Context) {
		typedObj := new(T)

		if err := c.BindJSON(&typedObj); err != nil {
			return
		}

		result := orm.Create(&typedObj)

		if result.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": result.Error.Error() })

			return
		}

		c.IndentedJSON(http.StatusCreated, &typedObj)
	}

	return handler
}
