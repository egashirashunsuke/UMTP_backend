package main

import (
	"log"
	"net/http"

	"github.com/egashirashunsuke/UMTP_backend/handler"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(".envファイルが見つかりません．")
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // React の開発サーバー
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		// 必要に応じて PUT, DELETE, OPTIONS なども追加
	}))

	//ルーティング
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.POST("/", handler.NewHintsHandler().GetHints)

	e.Logger.Fatal(e.Start(":8000"))
}
