package controller

import (
	"net/http"
	"strconv"

	"github.com/egashirashunsuke/UMTP_backend/dto"
	"github.com/egashirashunsuke/UMTP_backend/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IQuestionController interface {
	GetQuestionByID(c echo.Context) error
	GetAllQuestions(c echo.Context) error
	CreateQuestion(c echo.Context) error
	GetNextQuestion(c echo.Context) error
}

type questionController struct {
	uu usecase.IQuestionUsecase
}

func NewQuestionController(uu usecase.IQuestionUsecase) IQuestionController {
	return &questionController{uu: uu}
}

func (h *questionController) GetQuestionByID(c echo.Context) error {
	idStr := c.Param("questionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	question, err := h.uu.GetQuestionByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, question)
}

func (h *questionController) GetAllQuestions(c echo.Context) error {

	questions, err := h.uu.GetAllQuestions()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, questions)
}

func (h *questionController) CreateQuestion(c echo.Context) error {
	var in dto.CreateQuestionDTO
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := h.uu.CreateQuestion(&in); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "question created successfully"})
}

func (h *questionController) GetNextQuestion(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	nextQuestion, err := h.uu.GetNextQuestion(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nextQuestion)
}
