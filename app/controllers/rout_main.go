package controllers

import (
	"net/http"
	"shiba/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	shibaImgs := models.GetShibas("5", "true", "true")
	generateHTML(w, shibaImgs, "layout", "/top")
}
