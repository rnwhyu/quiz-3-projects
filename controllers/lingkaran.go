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

type LingkaranRequest struct {
	JariJari string `form:"jariJari" binding:"required,numeric"`
	Hitung   string `form:"hitung" binding:"required,oneof=keliling luas"`
}

func (pr LingkaranRequest) MapToLingkaran(lingkaran *models.Lingkaran) (err error) {
	jariJari, err := strconv.ParseFloat(pr.JariJari, 64)

	*lingkaran = models.Lingkaran{JariJari: jariJari}

	return
}
func HitungLingkaran(c *gin.Context) {
	var request LingkaranRequest

	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var lingkaran models.Lingkaran
	err = request.MapToLingkaran(&lingkaran)
	if err != nil {
		panic(err)
	}

	var result float64
	switch strings.ToLower(request.Hitung) {
	case "keliling":
		lingkaran.HitungKeliling()
		result = lingkaran.Keliling
	case "luas":
		lingkaran.HitungLuas()
		result = lingkaran.Luas
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf(
			"%s lingkaran dengan jari-jari %f adalah %f",
			cases.Title(language.English).String(request.Hitung),
			lingkaran.JariJari,
			result,
		),
		map[string]interface{}{
			"jari_jari":    lingkaran.JariJari,
			request.Hitung: result,
		},
	)
}