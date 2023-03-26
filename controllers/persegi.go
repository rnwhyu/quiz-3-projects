package controllers
import(
	"project-quiz3/response"
	"project-quiz3/models"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
type PersegiRequest struct {
	Sisi   string `form:"sisi" binding:"required,numeric"`
	Hitung string `form:"hitung" binding:"required,oneof=keliling luas"`
}

func (pr PersegiRequest) MapToPersegi(persegi *models.Persegi) (err error) {
	sisi, err := strconv.Atoi(pr.Sisi)

	*persegi = models.Persegi{Sisi: sisi}

	return
}
func HitungPersegi(c *gin.Context) {
	var request PersegiRequest

	err := c.ShouldBindQuery(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var persegi models.Persegi
	err = request.MapToPersegi(&persegi)
	if err != nil {
		panic(err)
	}

	var result int
	switch strings.ToLower(request.Hitung) {
	case "keliling":
		persegi.HitungKeliling()
		result = persegi.Keliling
	case "luas":
		persegi.HitungLuas()
		result = persegi.Luas
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf(
			"%s persegi dengan sisi %d adalah %d",
			cases.Title(language.English).String(request.Hitung),
			persegi.Sisi,
			result,
		),
		map[string]interface{}{
			"sisi":         persegi.Sisi,
			request.Hitung: result,
		},
	)
}