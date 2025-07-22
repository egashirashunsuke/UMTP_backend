package controller

import (
	"net/http"
	"strconv"

	"github.com/egashirashunsuke/UMTP_backend/usecase"

	"github.com/labstack/echo/v4"
)

type SubmitAnswerRequest struct {
	Answers map[string]*string `json:"answers"`
}

type IHintsController interface {
	GetHints(c echo.Context) error
}

type hintsController struct {
	uu usecase.IHintsUsecase
}

func NewHintsController(uu usecase.IHintsUsecase) IHintsController {
	return &hintsController{uu: uu}
}

func (h *hintsController) GetHints(c echo.Context) error {
	var req SubmitAnswerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストのバインドに失敗しました"})
	}

	idStr := c.Param("questionID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid questionID"})
	}

	out, err := h.uu.GetHints(c.Request().Context(),
		usecase.GenerateHintInput{
			QuestionID: id,
			Answers:    req.Answers,
		})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 3) 結果返却
	return c.JSON(http.StatusOK, out)
}
