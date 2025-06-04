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

	// service に質問を渡して結果を取得
	answer, err := h.hintsSvc.GetHints(c.Request().Context(), question)
	if err != nil {
		// OpenAI などへの問い合わせ中にエラーが起きた場合は 500 を返す
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 成功時は JSON で返却
	return c.JSON(http.StatusOK, map[string]string{"answer": answer})
}
