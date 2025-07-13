package handler

import (
	"net/http"
	"strconv"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type QuestionHandler struct {
	DB *gorm.DB
}

func NewQuestionHandler(db *gorm.DB) *QuestionHandler {
	return &QuestionHandler{DB: db}
}

func (h *QuestionHandler) GetQuestionByID(c echo.Context) error {
	idStr := c.Param("questionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	question, err := model.GetQuestionByID(h.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, question)
}

func (h *QuestionHandler) GetAllQuestions(c echo.Context) error {
	var questions []model.Question
	if err := h.DB.Preload("Choices").Find(&questions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, questions)
}

func (h *QuestionHandler) CreateQuestion(c echo.Context) error {
	var question model.Question
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := h.DB.Create(&question).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, question)
}
