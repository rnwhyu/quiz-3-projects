package controllers

import (
	"project-quiz3/models"
	"project-quiz3/response"
	"strconv"
	"fmt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
type SegitigaRequest struct {
	Alas   string `form:"alas" binding:"required,numeric"`
	Tinggi string`form:"tinggi" binding:"required,numeric"`
	Hitung string `form:"hitung" binding:"required,oneof=keliling luas"`
}

func (sr SegitigaRequest) MapToSegitiga(segitiga *models.Segitiga) (err error) {
	alas,err := strconv.ParseFloat(sr.Alas, 64)
	tinggi,err := strconv.ParseFloat(sr.Tinggi, 64)
	*segitiga = models.Segitiga{Alas:alas, Tinggi:tinggi}

	return
}
func HitungSegitiga(c *gin.Context) {
	var request SegitigaRequest

	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var segitiga models.Segitiga
	err = request.MapToSegitiga(&segitiga)
	if err != nil {
		panic(err)
	}

	var result float64
	switch strings.ToLower(request.Hitung) {
	case "keliling":
		segitiga.HitungKeliling()
		result = segitiga.Keliling
	case "luas":
		segitiga.HitungLuas()
		result = segitiga.Luas
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf(
			"%s segitiga sama sisi dengan sisi %f adalah %f",
			cases.Title(language.English).String(request.Hitung),
			segitiga.Alas,
			result,
		),
		map[string]interface{}{
			"sisi":         segitiga.Alas,
			request.Hitung: result,
		},
	)
}
