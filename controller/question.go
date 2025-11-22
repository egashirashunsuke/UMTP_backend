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
	GetPrevQuestion(c echo.Context) error
	CheckAnswer(c echo.Context) error
	GetAnswer(c echo.Context) error
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

func (h *questionController) GetPrevQuestion(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	prevQuestion, err := h.uu.GetPrevQuestion(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, prevQuestion)
}

func (h *questionController) CheckAnswer(c echo.Context) error {
	var req SubmitAnswerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストのバインドに失敗しました"})
	}

	idStr := c.Param("questionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid questionID"})
	}

	out, err := h.uu.CheckAnswer(id, req.Answers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"correct": out,
	})
}

func (h *questionController) GetAnswer(c echo.Context) error {
	idStr := c.Param("questionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid questionID"})
	}

	answer, err := h.uu.GetAnswer(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"answers": answer,
	})
}
