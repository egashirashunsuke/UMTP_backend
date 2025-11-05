package controller

import (
	"net/http"
	"time"

	"github.com/egashirashunsuke/UMTP_backend/dto"
	"github.com/egashirashunsuke/UMTP_backend/usecase"
	"github.com/labstack/echo/v4"
)

type ILogController interface {
	SendLog(c echo.Context) error
}

type LogController struct {
	uu usecase.ILogUsecase
}

func NewLogController(uu usecase.ILogUsecase) ILogController {
	return &LogController{
		uu: uu,
	}
}

func (lc *LogController) SendLog(c echo.Context) error {
	var in dto.LogRequest
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}

	var subPtr *string
	if v := c.Get("sub"); v != nil {
		if sub, ok := v.(string); ok && sub != "" {
			subPtr = &sub
		}
	}

	var clientAt *time.Time
	if in.Timestamp != nil {
		if t, err := time.Parse(time.RFC3339, *in.Timestamp); err == nil {
			clientAt = &t
		}
	}

	// Usecase入力DTOへ詰め替え（アプリ内DTO）
	uin := usecase.LogCommand{
		Sub:            subPtr,
		StudentNo6:     in.StudentID,
		QuestionID:     in.QuestionID,
		EventName:      in.EventName,
		Answers:        in.Answers,
		HintOpenStatus: in.HintOpenStatus,
		Hints:          in.Hints,
		HintIndex:      in.HintIndex,
		Useful:         in.Useful,
		Comment:        in.Comment,
		AnonID:         in.AnonID,
		ClientAt:       clientAt,
	}

	id, err := lc.uu.SendLog(c.Request().Context(), uin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save log"})
	}
	return c.JSON(http.StatusCreated, map[string]any{"id": id})
}
