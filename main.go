package main

import (
	"net/http"

	"github.com/egashirashunsuke/UMTP_backend/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // React の開発サーバー
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		// 必要に応じて PUT, DELETE, OPTIONS なども追加
	}))
	e.GET("/", handler.CreateHintsHandler())

	e.Logger.Fatal(e.Start(":8000"))
}
