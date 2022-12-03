package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/db"
	"gorm.io/gorm"
)

func GetByIdH[T any](orm *gorm.DB, args ...string) func (*gin.Context) {
	var preload string
	
	if len(args) > 0 {
		preload = args[0]
	}


	handler := func (c *gin.Context) {
		var itemParams db.IDParams

		if err := c.ShouldBindUri(&itemParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })

			return
		}
		
		var item T

		tx := orm.Model(new(T)).Where("id = ?", itemParams.ID).Preload(preload).First(&item)

		if tx.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": tx.Error.Error() })

			return
		}

		c.IndentedJSON(http.StatusOK, &item)
	}

	return handler
}
