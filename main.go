package main

import (
	"log"
	"net/http"
	"os"

	"github.com/egashirashunsuke/UMTP_backend/handler"
	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db := model.DBConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := godotenv.Load(); err != nil {
		log.Println(".envファイルが見つかりません．")
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://umtp-shunsuke-egashiras-projects.vercel.app",
		},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		// 必要に応じて PUT, DELETE, OPTIONS なども追加
	}))

	//ルーティング
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.POST("/", handler.NewHintsHandler().GetHints)

	e.GET("/quesion/:questionID", handler.NewQuestionHandler(db).GetQuestionByID)

	// PORT環境変数を取得し、なければ10000を使う
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
