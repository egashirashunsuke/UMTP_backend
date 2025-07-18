package model

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// DB接続とテーブルを作成する
func DBConnection() *gorm.DB {
	dsn := GetDBConfig()
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	CreateTable(db)
	return db
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

func CreateTable(db *gorm.DB) {
	// テーブル作成
	db.AutoMigrate(&Question{}, &Choice{}, &Label{}, &AnswerMapping{})

	// サンプルQuestion
	question := Question{
		ProblemDescription: `
市営駐輪場の利用申し込みシステムです。
この市には第1と第2の駐輪場があります。それぞれの駐輪場で、月額料金は異なります。
市営駐輪場の利用を申し込むには、事前に会員登録する必要があります。会員登録では、氏名、住所、電話番号、メールアドレスを登録します。会員登録をすると、登録番号とパスワードがメールで送付されます。駐輪場の申し込み画面を開き、登録番号とパスワードを入力し、ログインしたあとに、希望の駐輪場を選択し利用申請をします。希望者が多数の場合は、抽選が行われます。抽選が外れると、次回の抽選で当選しやすくなります。`,
		Question:          "次のクラス図のa〜gに該当するものを選択し、モデルを完成させなさい。",
		ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram.png?alt=media&token=76ecd79f-1317-42a5-bcc9-c31278982abc",
		ClassDiagramPlantUML: `@startuml 
class a 
class b 
class c 
class d 
class e 
class f 
class g
a o-- b 
b -- c 
c -- d 
c -- e 
e <|-- f 
e <|-- g 
@enduml
`,
	}

	// QuestionをFirstOrCreate
	if err := db.Where("question = ?", question.Question).FirstOrCreate(&question).Error; err != nil {
		panic(err)
	}

	// Choices
	choices := []Choice{
		{ChoiceCode: "A", ChoiceText: "駐輪場", QuestionID: question.ID},
		{ChoiceCode: "B", ChoiceText: "第1駐輪場", QuestionID: question.ID},
		{ChoiceCode: "C", ChoiceText: "第2駐輪場", QuestionID: question.ID},
		{ChoiceCode: "D", ChoiceText: "会員", QuestionID: question.ID},
		{ChoiceCode: "E", ChoiceText: "申請", QuestionID: question.ID},
		{ChoiceCode: "F", ChoiceText: "抽選結果", QuestionID: question.ID},
		{ChoiceCode: "G", ChoiceText: "会員リスト", QuestionID: question.ID},
	}
	for _, c := range choices {
		db.Where("choice_code = ? AND question_id = ?", c.ChoiceCode, question.ID).FirstOrCreate(&c)
	}

	// Labels
	labels := []Label{
		{LabelCode: "a", QuestionID: question.ID, OrderId: 1},
		{LabelCode: "b", QuestionID: question.ID, OrderId: 2},
		{LabelCode: "c", QuestionID: question.ID, OrderId: 3},
		{LabelCode: "d", QuestionID: question.ID, OrderId: 4},
		{LabelCode: "e", QuestionID: question.ID, OrderId: 5},
		{LabelCode: "f", QuestionID: question.ID, OrderId: 6},
		{LabelCode: "g", QuestionID: question.ID, OrderId: 7},
	}
	for _, l := range labels {
		db.Where("label_code = ? AND question_id = ?", l.LabelCode, question.ID).FirstOrCreate(&l)
	}

	// AnswerMapping例（A-a, B-b, ...）
	mappings := []struct {
		LabelCode  string
		ChoiceCode string
	}{
		{"a", "G"},
		{"b", "D"},
		{"c", "E"},
		{"d", "F"},
		{"e", "A"},
		{"f", "B"},
		{"g", "C"},
	}

	for _, m := range mappings {
		// Choice を取得
		var c Choice
		if err := db.
			Where("choice_code = ? AND question_id = ?", m.ChoiceCode, question.ID).
			First(&c).Error; err != nil {
			// 見つからなければスキップ or ログ出力
			continue
		}

		// Label を取得
		var l Label
		if err := db.
			Where("label_code = ? AND question_id = ?", m.LabelCode, question.ID).
			First(&l).Error; err != nil {
			continue
		}

		// マッピングを作成 or 既存取得
		am := AnswerMapping{
			QuestionID: question.ID,
			ChoiceID:   c.ID,
			LabelID:    l.ID,
		}
		db.
			Where("question_id = ? AND choice_id = ? AND label_id = ?", am.QuestionID, am.ChoiceID, am.LabelID).
			FirstOrCreate(&am)
	}
}
