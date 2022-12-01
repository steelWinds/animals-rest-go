package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetH[T any](orm *gorm.DB, args ...string) func (*gin.Context) {
	var preload string
	
	if len(args) > 0 {
		preload = args[0]
	}

	handler := func (c *gin.Context) {
		items := make([]T, 0)

		var tx *gorm.DB

		if len(preload) > 0 {
			tx = orm.Model(new(T)).Preload(preload).Find(&items)
		} else {
			tx = orm.Find(&items)
		}

		if tx.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": "Empty list" })
		}

		c.IndentedJSON(http.StatusOK, &items)
	}

	return handler
}
