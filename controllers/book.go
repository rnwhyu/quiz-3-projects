package controllers

import (
	"fmt"
	"net/http"
	"project-quiz3/models"
	"project-quiz3/response"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookCreateRequest struct {
	Category_id  int    `json:"category_id" binding:"required,numeric"`
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Image_url    string `json:"image_url" binding:"required"`
	Release_year int    `json:"release_year" binding:"required,numeric"`
	Price        string `json:"price" binding:"required"`
	Total_page   string `json:"total_page" binding:"required,numeric"`
}

func (bcr BookCreateRequest) MapToBook(book *models.Book) error {
	totalPage, err := strconv.Atoi(bcr.Total_page)
	if err != nil {
		return err
	}

	thickness := ""
	if totalPage > 200 {
		thickness = "tebal"
	} else if totalPage > 100 {
		thickness = "sedang"
	} else {
		thickness = "tipis"
	}

	*book = models.Book{
		Category_id:  bcr.Category_id,
		Title:        bcr.Title,
		Description:  bcr.Description,
		Image_url:    bcr.Image_url,
		Release_year: bcr.Release_year,
		Price:        bcr.Price,
		Total_page:   totalPage,
		Thickness:    thickness,
	}

	return nil
}

type BookFindRequest struct {
	ID string `uri:"id" binding:"required,numeric"`
}

type BookUpdateRequest struct {
	BookFindRequest
	BookCreateRequest
}

func FindAllBooks(c *gin.Context) {
	books := models.Books{}
	err := books.FindAll()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(c, http.StatusOK, "Success", books)
}
func CreateBook(c *gin.Context) {
	var request BookCreateRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	book := models.Book{}
	err = request.MapToBook(&book)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var validationErrors []map[string]string
	validUrl, _ := regexp.MatchString(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`+regexp.QuoteMeta(request.Image_url)+`(?:/[^/\s]+)*/?$`, book.Image_url)
	if !validUrl {
		validationErrors = append(validationErrors, map[string]string{
			"field":   "image_url",
			"message": "Invalid URL",
		})
	}

	if book.Release_year < 1980 || book.Release_year > 2021 {
		validationErrors = append(validationErrors, map[string]string{
			"field":   "release_year",
			"message": "Must be between 1980 and 2021",
		})
	}

	if len(validationErrors) > 0 {
		response.ApiErrorResponse(c, http.StatusBadRequest, "Validation Error", validationErrors)
		return
	}

	err = book.Create()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(
		c,
		http.StatusCreated,
		fmt.Sprintf("Book #%d created successfully", book.ID),
		book,
	)
}

func UpdateBook(c *gin.Context) {
	var bookId BookFindRequest
	var request BookUpdateRequest

	err := c.ShouldBindUri(&bookId)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	request.ID = bookId.ID

	err = c.ShouldBindJSON(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(request.ID)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	book := models.Book{}
	err = request.MapToBook(&book)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var validationErrors []map[string]string
	validUrl, _ := regexp.MatchString(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`+regexp.QuoteMeta(request.Image_url)+`(?:/[^/\s]+)*/?$`, book.Image_url)
	if !validUrl {
		validationErrors = append(validationErrors, map[string]string{
			"field":   "image_url",
			"message": "Invalid URL",
		})
	}

	if book.Release_year < 1980 || book.Release_year > 2021 {
		validationErrors = append(validationErrors, map[string]string{
			"field":   "release_year",
			"message": "Must be between 1980 and 2021",
		})
	}

	if len(validationErrors) > 0 {
		response.ApiErrorResponse(c, http.StatusBadRequest, "Validation Error", validationErrors)
		return
	}

	book.ID = id

	err = book.Update()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf("Book #%d updated successfully", book.ID),
		book,
	)
}

func DeleteBook(c *gin.Context) {
	var request BookFindRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(request.ID)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	book := models.Book{ID: id}
	err = book.Delete()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf("Book #%d deleted successfully", book.ID),
	)
}
