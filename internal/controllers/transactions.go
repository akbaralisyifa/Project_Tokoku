package controllers

import (
	"tokoku/internal/models"
)

type TransController struct {
	model *models.TransModel
}

func NewTransController(m *models.TransModel) *TransController {
	return &TransController{
		model: m,
	}
}
