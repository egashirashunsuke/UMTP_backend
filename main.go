package main

import (
	"log"
	"net/http"
	"os"

	"github.com/egashirashunsuke/UMTP_backend/controller"
	"github.com/egashirashunsuke/UMTP_backend/model"
	"github.com/egashirashunsuke/UMTP_backend/repository"
	"github.com/egashirashunsuke/UMTP_backend/usecase"
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

	qRepo := repository.NewQuestionRepository(db)
	qUsecase := usecase.NewQuestionUsecase(qRepo)
	qCtrl := controller.NewQuestionController(qUsecase)

	//ルーティング
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/question/:questionID", qCtrl.GetQuestionByID)
	e.GET("/questions", qCtrl.GetAllQuestions)

	// e.POST("/:questionID", handler.NewHintsHandler(db).GetHints)
	// e.GET("/questions", handler.NewQuestionHandler(db).GetAllQuestions)
	// e.POST("/question", handler.NewQuestionHandler(db).CreateQuestion)

	// PORT環境変数を取得し、なければ10000を使う
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
