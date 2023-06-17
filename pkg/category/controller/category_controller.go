package controller

import (
	"final-project3/pkg/category/dto"
	"final-project3/pkg/category/model"
	"final-project3/pkg/category/usecase"
	"final-project3/utils/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHTTPController struct {
	usecase usecase.UsecaseInterfaceCategory
}

func InitControllerCategory(uc usecase.UsecaseInterfaceCategory) *CategoryHTTPController {
	return &CategoryHTTPController{
		usecase: uc,
	}
}

func (uc *CategoryHTTPController) CreateNewCategory(c *gin.Context) {
	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors := helpers.FormatError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	category, err := uc.usecase.CreateNewCategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         category.Id,
		"type":       category.Type,
		"created_at": category.CreatedAt,
	})
}

func (uc *CategoryHTTPController) GetAllCategory(c *gin.Context) {
	categories, err := uc.usecase.GetAllCategory()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res []dto.CategoryResponse
	for _, category := range categories {
		temp := ConvertToCategoryResponse(category)
		res = append(res, temp)
	}

	c.JSON(http.StatusOK, res)
}

func (uc *CategoryHTTPController) UpdateCategoryById(c *gin.Context) {
	idString := c.Param("categoryId")
	categoryId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing Category ID", "status": http.StatusBadRequest})
		return
	}
	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors := helpers.FormatError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	category, err := uc.usecase.UpdateCategoryById(categoryId, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         category.Id,
		"type":       category.Type,
		"updated_at": category.UpdatedAt,
	})
}

func (uc *CategoryHTTPController) DeleteCategoryById(c *gin.Context) {

	idString := c.Param("categoryId")
	categoryId, err := strconv.Atoi(idString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Parsing Category ID", "status": http.StatusBadRequest})
		return
	}
	err = uc.usecase.DeleteCategoryById(categoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}

func ConvertToCategoryResponse(category model.Category) dto.CategoryResponse {
	var tasks []dto.TaskData
	for _, task := range category.Tasks {
		tasks = append(tasks, dto.TaskData{
			Id:          task.Id,
			Title:       task.Title,
			Description: task.Description,
			UserId:      task.UserId,
			CategoryId:  task.CategoryId,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return dto.CategoryResponse{
		Id:        category.Id,
		Type:      category.Type,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Tasks:     tasks,
	}
}
