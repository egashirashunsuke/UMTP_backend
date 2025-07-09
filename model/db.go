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

// テーブルを作成する
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Question{}, &Choice{})
	question := Question{
		ProblemDescription: `
市営駐輪場の利用申し込みシステムです。
この市には第1と第2の駐輪場があります。それぞれの駐輪場で、月額料金は異なります。
市営駐輪場の利用を申し込むには、事前に会員登録する必要があります。会員登録では、氏名、住所、電話番号、メールアドレスを登録します。会員登録をすると、登録番号とパスワードがメールで送付されます。駐輪場の申し込み画面を開き、登録番号とパスワードを入力し、ログインしたあとに、希望の駐輪場を選択し利用申請をします。希望者が多数の場合は、抽選が行われます。抽選が外れると、次回の抽選で当選しやすくなります。`,
		Question: `次のクラス図のa〜gに該当するものを選択し、モデルを完成させなさい。`,
		Image: "classdiagram.png",
		Choices: []Choice{
			{Label: "A", Text: "駐輪場"},
			{Label: "B", Text: "第1駐輪場"},
			{Label: "C", Text: "第2駐輪場"},
			{Label: "D", Text: "会員"},
			{Label: "E", Text: "申請"},
			{Label: "F", Text: "抽選結果"},
			{Label: "G", Text: "会員リスト"},
		},
	}
	db.
		Where("question = ?", question.Question).
		FirstOrCreate(&question)

	question2 := Question{
		ProblemDescription: `
経営診断システムの開発を検討しています。
診断する会社についての以下の情報を入力することで、会社の経営診断結果が表示されます。
・基本データ（業種、従業員数など）
・財務データ（流動資産、固定資産、売上、利益など）
診断は、業種別に蓄積されている既存の蓄積データと比較することにより行われます。
今回診断のため入力した会社データも、この蓄積データに追加されます。
診断結果は、総合評価と、収益性、効率性、安全性、成長性のそれぞれの診断指標で表示されます。
診断は、1社1回のみ行うことができます。`,
		Question: `次のクラス図のa~gに該当するものを選択しなさい。
なお、該当する選択肢が複数ある場合は、選択肢のアルファベットの早い順に選択しなさい。`,
		Image: "classdiagram2.png",
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
		Where("question = ?", question2.Question).
		FirstOrCreate(&question2)

	question3 := Question{
		ProblemDescription: `
ある国では、電子入国許可システムを採用しています。 
申請者が、Web上から電子ビザの登録を申請すると、参照番号が発行されます。
登録には、パスポートの情報およびクレジットカードの情報が必要になります。
電子ビザには、観光用と短期商用のものがあります。観光用のものは、有効期限内ならば何回でも使用できますが、短期商用のものは1回限りです。`,
		Question: `次のクラス図のa~gに該当するものを選択しなさい。`,
		Image: "classdiagram3.png",
		Choices: []Choice{
			{Label: "A", Text: "パスポート"},
			{Label: "B", Text: "電子ビザ"},
			{Label: "C", Text: "申請登録"},
			{Label: "D", Text: "申請者"},
			{Label: "E", Text: "履歴"},
			{Label: "F", Text: "観光用電子ビザ"},
			{Label: "G", Text: "短期商用電子ビザ"},
		},
	}
	db.
		Where("question = ?", question3.Question).
		FirstOrCreate(&question3)
}
