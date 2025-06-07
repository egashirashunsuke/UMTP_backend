package handler

import (
	"net/http"

	"github.com/egashirashunsuke/UMTP_backend/service"

	"github.com/labstack/echo/v4"
)

type SubmitAnswerRequest struct {
	Answers map[string]*string `json:"answers"`
}

type HintsHandler struct {
	hintsSvc service.HintsService
}

func NewHintsHandler() *HintsHandler {
	return &HintsHandler{
		hintsSvc: service.NewHintsService(),
	}
}

func (h *HintsHandler) GetHints(c echo.Context) error {
	var req SubmitAnswerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストのバインドに失敗しました"})

	}

	answer, err := h.hintsSvc.GetHints(c.Request().Context(), req.Answers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": answer})
}
