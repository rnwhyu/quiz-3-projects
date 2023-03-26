package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"project-quiz3/models"
	"project-quiz3/response"
	"strconv"
	"strings"
)
type PersegiPanjangRequest struct {
	Panjang string `form:"panjang" binding:"required,numeric"`
	Lebar   string `form:"lebar" binding:"required,numeric"`
	Hitung  string `form:"hitung" binding:"required,oneof=keliling luas"`
}

func (pp PersegiPanjangRequest) MapToPersegiPanjang(persegiPanjang *models.PersegiPanjang) error {
	panjang, err := strconv.Atoi(pp.Panjang)
	if err != nil {
		return err
	}

	lebar, err := strconv.Atoi(pp.Lebar)
	if err != nil {
		return err
	}

	*persegiPanjang = models.PersegiPanjang{Panjang: panjang, Lebar: lebar}

	return nil
}
func HitungPP(c *gin.Context) {
	var request PersegiPanjangRequest

	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var pp models.PersegiPanjang
	err = request.MapToPersegiPanjang(&pp)
	if err != nil {
		panic(err)
	}

	var result int
	switch strings.ToLower(request.Hitung) {
	case "keliling":
		pp.HitungKeliling()
		result = pp.Keliling
	case "luas":
		pp.HitungLuas()
		result = pp.Luas
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf(
			"%s persegi panjang dengan panjang %d dan lebar %d adalah %d",
			cases.Title(language.English).String(request.Hitung),
			pp.Panjang,
			pp.Lebar,
			result,
		),
		map[string]interface{}{
			"panjang":      pp.Panjang,
			"lebar":        pp.Lebar,
			request.Hitung: result,
		},
	)
}