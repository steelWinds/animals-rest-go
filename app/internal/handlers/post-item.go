package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelWinds/animals-rest-go/internal/database"
)

func PostItem[T any](c *gin.Context) {
	item := new(T)

	if err := c.BindJSON(&item); err != nil {
		return
	}

	orm := database.GetDB()

	if tx := orm.Create(&item); tx.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ "message": tx.Error.Error() })

		return
	}

	c.IndentedJSON(http.StatusCreated, &item)
}
