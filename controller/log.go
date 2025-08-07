package controller

import (
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
	logData := new(usecase.LogData)
	if err := c.Bind(logData); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid log data"})
	}

	if err := lc.uu.SendLog(logData); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to process log"})
	}

	return c.JSON(200, map[string]string{"message": "Log processed successfully"})
}
