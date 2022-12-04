package models

type IDParam struct {
	ID uint `uri:"id" binding:"required"`
}