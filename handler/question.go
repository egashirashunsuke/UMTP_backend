package handler

import (
	"net/http"
	"strconv"

	"github.com/egashirashunsuke/UMTP_backend/dto"
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
	var q dto.CreateQuestionDTO
	if err := c.Bind(&q); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	// モデル層に丸投げ
	if err := model.CreateQuestionWithAssociations(h.DB, &q); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 保存後の q.ID, CreatedAt, Choices[].ID などもセット済み
	return c.JSON(http.StatusCreated, q)
}
