package model

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// DB接続とテーブルを作成する
func DBConnection() *sql.DB {
	dsn := GetDBConfig()
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	CreateTable(db)
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	return sqlDB
}

// DBのdsnを取得する
func GetDBConfig() string {

	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port,
	)
	return dsn
}

// テーブルを作成する
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Question{}, &Choice{})
	question := Question{
		Question: `次の問題記述を読んで、設問に答えなさい。

## 問題記述:
経営診断システムの開発を検討しています。
診断する会社についての以下の情報を入力することで、会社の経営診断結果が表示されます。
・基本データ（業種、従業員数など）
・財務データ（流動資産、固定資産、売上、利益など）

診断は、業種別に蓄積されている既存の蓄積データと比較することにより行われます。
今回診断のため入力した会社データも、この蓄積データに追加されます。
診断結果は、総合評価と、収益性、効率性、安全性、成長性のそれぞれの診断指標で表示されます。
診断は、1社1回のみ行うことができます。

## 設問1:
次のクラス図のa~gに該当するものを選択しなさい。
なお、該当する選択肢が複数ある場合は、選択肢のアルファベットの早い順に選択しなさい。`,
		Choices: []Choice{
			{Label: "A", Text: "診断指標"},
			{Label: "B", Text: "基本データ"},
			{Label: "C", Text: "診断結果"},
			{Label: "D", Text: "会社情報"},
			{Label: "E", Text: "財務データ"},
			{Label: "F", Text: "自社情報"},
			{Label: "G", Text: "業種"},
		},
	}
	db.
		Where("question = ?", question.Question).
		FirstOrCreate(&question)
}
