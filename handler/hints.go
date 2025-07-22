package handler

import (
	"net/http"
	"strconv"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/service"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type SubmitAnswerRequest struct {
	Answers map[string]*string `json:"answers"`
}

type HintsHandler struct {
	hintsSvc service.HintsService
	DB       *gorm.DB
}

func NewHintsHandler(db *gorm.DB) *HintsHandler {
	return &HintsHandler{
		hintsSvc: service.NewHintsService(),
		DB:       db,
	}
}

func (h *HintsHandler) GetHints(c echo.Context) error {
	var req SubmitAnswerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストのバインドに失敗しました"})
	}

	idStr := c.Param("questionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid questionID"})
	}

	question, err := model.GetQuestionByID(h.DB, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "question not found"})
	}

	answer, err := h.hintsSvc.GetHints(c.Request().Context(), question, req.Answers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": answer})
}
