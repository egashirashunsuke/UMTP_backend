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

	origin := c.Request().Header.Get("Origin")
	if origin == "" {
		origin = c.Request().Header.Get("Referer")
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

	uin := usecase.LogCommand{
		Sub:            subPtr,
		QuestionID:     in.QuestionID,
		EventName:      in.EventName,
		Answers:        in.Answers,
		HintOpenStatus: in.HintOpenStatus,
		Hints:          in.Hints,
		AnonID:         in.AnonID,
		ClientOrigin:   &origin,
		ClientAt:       clientAt,
	}

	id, err := lc.uu.SendLog(c.Request().Context(), uin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save log"})
	}
	return c.JSON(http.StatusCreated, map[string]any{"id": id})
}
