package controllers

import (
	"fmt"
	"net/http"
	"project-quiz3/models"
	"project-quiz3/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryFindRequest struct {
	ID string `uri:"id" binding:"required,numeric"`
}

type CategoryUpdateRequest struct {
	CategoryFindRequest
	CategoryCreateRequest
}

func FindAll(c *gin.Context) {
	categories := models.Categories{}
	err := categories.FindAllCategory()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(c, http.StatusOK, "Success", categories)
}

func FindBooks(c *gin.Context) {
	var request CategoryFindRequest

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

	books := models.Books{}
	err = books.FindByCategoryID(id)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(c, http.StatusOK, "Success", books)
}

func CreateCategory(c *gin.Context) {
	var request CategoryCreateRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category := models.Category{Name: request.Name}
	err = category.CreateCategory()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(
		c,
		http.StatusCreated,
		fmt.Sprintf("Category #%d created successfully", category.ID),
		category,
	)
}

func FindByID(c *gin.Context) {
	var request CategoryFindRequest

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

	category := models.Category{ID: id}
	err = category.FindCategory()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		"Success",
		category,
	)
}

func UpdateCategory(c *gin.Context) {
	var requestId CategoryFindRequest
	var request CategoryUpdateRequest

	err := c.ShouldBindUri(&requestId)
	if err != nil {
		response.ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	request.ID = requestId.ID

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

	category := models.Category{ID: id, Name: request.Name}
	err = category.UpdateCategory()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf("Category #%d updated successfully", category.ID),
		category,
	)
}

func DeleteCategory(c *gin.Context) {
	var request CategoryFindRequest

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

	category := models.Category{ID: id}
	err = category.DeleteCategory()
	if err != nil {
		response.ApiErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.ApiResponse(
		c,
		http.StatusOK,
		fmt.Sprintf("Category #%d updated successfully", category.ID),
	)
}
