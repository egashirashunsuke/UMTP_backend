package handler

import (
	"net/http"

	"github.com/egashirashunsuke/UMTP_backend/service"

	"github.com/labstack/echo/v4"
)

type HintsHandler struct {
	hintsSvc service.HintsService
}

func NewHintsHandler() *HintsHandler {
	return &HintsHandler{
		hintsSvc: service.NewHintsService(),
	}
}

func (h *HintsHandler) GetHints(c echo.Context) error {
	question := c.QueryParam("question")
	if question == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "質問をクエリパラメータ 'question' に指定してください"})
	}

	answer, err := h.hintsSvc.GetHints(c.Request().Context(), question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"answer": answer})
}
