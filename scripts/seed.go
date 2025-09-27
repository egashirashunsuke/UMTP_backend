package main

import (
	"log"

	"github.com/egashirashunsuke/UMTP_backend/model"
	"gorm.io/gorm"
)

// seed 用の構造体
type QuestionSeed struct {
	Question model.Question
	Choices  []model.Choice
	Labels   []model.Label
	Mappings []struct {
		LabelCode  string
		ChoiceCode string
	}
}

func main() {
	db := model.DBConnection()
	if err := seedQuestions(db); err != nil {
		log.Fatalf("seed failed: %v", err)
	}
	log.Println("✅ Seeding complete")
}

func seedQuestions(db *gorm.DB) error {
	seeds := []QuestionSeed{
		{
			Question: model.Question{
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
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "駐輪場"},
				{ChoiceCode: "B", ChoiceText: "第1駐輪場"},
				{ChoiceCode: "C", ChoiceText: "第2駐輪場"},
				{ChoiceCode: "D", ChoiceText: "会員"},
				{ChoiceCode: "E", ChoiceText: "申請"},
				{ChoiceCode: "F", ChoiceText: "抽選結果"},
				{ChoiceCode: "G", ChoiceText: "会員リスト"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
				{LabelCode: "g", OrderId: 7},
			},
			Mappings: []struct {
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
			},
		},
		{
			Question: model.Question{
				ProblemDescription: `
経営診断システムの開発を検討しています。
診断する会社についての以下の情報を入力することで、会社の経営診断結果が表示されます。
・基本データ（業種、従業員数など）
・財務データ（流動資産、固定資産、売上、利益など）
診断は、業種別に蓄積されている既存の蓄積データと比較することにより行われます。今回診断のため入力した会社データも、この蓄積データに追加されます。診断結果は、総合評価と、収益性、効率性、安全性、成長性のそれぞれの診断指標で表示されます。診断は、1社1回のみ行うことができます。`,
				Question:          "次のクラス図のa~gに該当するものを選択しなさい。なお、該当する選択肢が複数ある場合は、選択肢のアルファベットの早い順に選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram2.png?alt=media&token=7555a4dd-7e75-4d11-ae95-47a37aecbc53",
				ClassDiagramPlantUML: `@startuml
class a { }
class b { }
class c { }
class d { }
class e { }
class f { }
class g { }
蓄積データ "1" o-- "0.." a
a "1" o-- "0.." b
b "1" o-- "h" c
b "1" o-- "i" d
c "j" --o "0..1" g
d "k" --o "0..1" g
e "l" --o "1" f
f "1" --o "1" g
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "診断指標"},
				{ChoiceCode: "B", ChoiceText: "基本データ"},
				{ChoiceCode: "C", ChoiceText: "診断結果"},
				{ChoiceCode: "D", ChoiceText: "会社情報"},
				{ChoiceCode: "E", ChoiceText: "財務データ"},
				{ChoiceCode: "F", ChoiceText: "自社情報"},
				{ChoiceCode: "G", ChoiceText: "業種"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
				{LabelCode: "g", OrderId: 7},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "G"},
				{"b", "D"},
				{"c", "B"},
				{"d", "E"},
				{"e", "A"},
				{"f", "C"},
				{"g", "F"},
			},
		},
		{
			Question: model.Question{
				ProblemDescription: `
ある国では、電子入国許可システムを採用しています。
申請者が、Web上から電子ビザの登録を申請すると、参照番号が発行されます。
登録には、パスポートの情報およびクレジットカードの情報が必要になります。
電子ビザには、観光用と短期商用のものがあります。観光用のものは、有効期限内ならば何回でも使用できますが、短期商用のものは1回限りです。`,
				Question:          "次のクラス図のa~gに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram3.png?alt=media&token=79667b65-8f9e-483b-8b87-12ee1d943a3f",
				ClassDiagramPlantUML: `@startuml
class b {
i
参照番号
}
class a {
h
}
class d
class e
class g
class c {
j
k
}
class f {
l
有効期限
}
class "クレジットカード" {
クレジット番号
有効期限
}
b "1" -- "0..1" a
a <|-- d
a <|-- e
b "1.." -- "1" c
d "1" o-- "0.." g
c o-- f
c o-- "クレジットカード"
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "パスポート"},
				{ChoiceCode: "B", ChoiceText: "電子ビザ"},
				{ChoiceCode: "C", ChoiceText: "申請登録"},
				{ChoiceCode: "D", ChoiceText: "申請者"},
				{ChoiceCode: "E", ChoiceText: "履歴"},
				{ChoiceCode: "F", ChoiceText: "観光用ビザ"},
				{ChoiceCode: "G", ChoiceText: "短期商用電子ビザ"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
				{LabelCode: "g", OrderId: 7},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "B"},
				{"b", "C"},
				{"c", "D"},
				{"d", "F"},
				{"e", "G"},
				{"f", "A"},
				{"g", "E"},
			},
		},
		{
			Question: model.Question{
				ProblemDescription: `
家庭教師の検索システムです。生徒は、システムに次のような家庭教師に対する希望（希望家庭教師像）を1〜3つまで入力します。
希望に対応した、家庭教師の候補が表示されます。
希望エリア（通常は自宅と同じ）：都道府県、市区町村群
希望性別：男/女/どちらでもよい
希望職業：学生/社会人／どちらでもよい
学校種別：小学校/中学校/高校
小学校：数学、国語、理科、社会
中学校：数学、国語、理科、社会、英語
高校：数学、国語、理科系、社会系、英語`,
				Question:          "クラス図のa〜eに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram2.png?alt=media&token=7555a4dd-7e75-4d11-ae95-47a37aecbc53",
				ClassDiagramPlantUML: `@startuml
class 生徒 {
氏名
住所
学校種別
}
class a
class b
class c
class d
class e
a "f" --o "1" 生徒
a "g" o-- "1" b
b "1" o-- "h" d
d "k" -- "j" e
c "1" o-- "i" e
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "家庭教師"},
				{ChoiceCode: "B", ChoiceText: "科目"},
				{ChoiceCode: "C", ChoiceText: "エリア"},
				{ChoiceCode: "D", ChoiceText: "希望家庭教師像"},
				{ChoiceCode: "E", ChoiceText: "学校"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "D"},
				{"b", "C"},
				{"c", "E"},
				{"d", "A"},
				{"e", "B"},
			},
		},
		{
			Question: model.Question{
				ProblemDescription: `
感染症は、病原体が原因で発生します。病原体には、寄生虫、細菌、真菌、ウイルスがあります。
感染経路には、飛沫感染、空気感染、接触感染、経口感染があります。
治療法には、薬剤によるものと外科的治療があります。また治療法のほか、予防法を検討する必要があります。
すべての感染症に関して、感染経路と病原体は判明していますが、治療法と予防法については、ないものもあります。`,
				Question:          "次のクラス図のa〜gに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram2.png?alt=media&token=7555a4dd-7e75-4d11-ae95-47a37aecbc53",
				ClassDiagramPlantUML: `@startuml
class a
class b
class c
class d
class e
class f
class g
class 病原体
class 外科的治療
a <|-- d
a <|-- 外科的治療
a "h" -- "i" b
b "j" o-- "k" c
b "l" -- "m" f
b "n" -- "o" 病原体
病原体 <|-- g
c <|-- e
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "飛沫感染 空気感染 接触感染 経口感染"},
				{ChoiceCode: "B", ChoiceText: "寄生虫 細菌 真菌 ウイルス"},
				{ChoiceCode: "C", ChoiceText: "治療法"},
				{ChoiceCode: "D", ChoiceText: "感染症"},
				{ChoiceCode: "E", ChoiceText: "感染経路"},
				{ChoiceCode: "F", ChoiceText: "薬剤"},
				{ChoiceCode: "G", ChoiceText: "予防法"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
				{LabelCode: "g", OrderId: 7},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "C"},
				{"b", "D"},
				{"c", "E"},
				{"d", "F"},
				{"e", "A"},
				{"f", "G"},
				{"g", "B"},
			},
		},
		{
			Question: model.Question{
				ProblemDescription: `
ある自転車店のパンク修理の費用を考えます。
パンク修理の費用は工賃と部品代の合計になります。
パンク修理は虫ゴムの交換のみ、パンク修理用パッチ処理、チューブ交換、チューブとタイヤの交換があります。虫ゴムの交換はどのパンク修理でも必ず行います。また、チューブだけの交換とチューブとタイヤの交換では工賃は変わりませんが、前輪と後輪では工賃は変わります。
後輪の場合は、後輪のブレーキ調整が必要です。またギア付き自転車の場合、後輪ギア調整も必要です。
修理は、前輪か後輪のどちらか1か所とします。`,
				Question:          "次のクラス図のa〜iに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram2.png?alt=media&token=7555a4dd-7e75-4d11-ae95-47a37aecbc53",
				ClassDiagramPlantUML: `@startuml
class a {
}
class b {
}
class c {
}
class d {
}
class e {
}
class f {
}
class g {
}
class h {
}
class i {
}
class タイヤ {
}
class 虫ゴム {
}
class パンク修理用パッチ {
}
class チューブ {
}
class 虫ゴム交換 {
}
a "1" o-- "n" b
b <|-- タイヤ
b <|-- 虫ゴム
b <|-- パンク修理用パッチ
b <|-- チューブ
a "1" o-- "j" c
c "1" o-- "k" 虫ゴム交換
c "1" o-- "l" d
c "1" o-- "m" e
e <|-- f
e <|-- g
g "0..1" o-- "0..1" h
g "0..1" o-- "1" i
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "部品"},
				{ChoiceCode: "B", ChoiceText: "後輪着脱"},
				{ChoiceCode: "C", ChoiceText: "前輪着脱"},
				{ChoiceCode: "D", ChoiceText: "車輪着脱"},
				{ChoiceCode: "E", ChoiceText: "パンク修理"},
				{ChoiceCode: "F", ChoiceText: "パンク修理工賃"},
				{ChoiceCode: "G", ChoiceText: "パッチ処理"},
				{ChoiceCode: "H", ChoiceText: "後輪ギア調整"},
				{ChoiceCode: "I", ChoiceText: "後輪ブレーキ調整"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
				{LabelCode: "g", OrderId: 7},
				{LabelCode: "h", OrderId: 8},
				{LabelCode: "i", OrderId: 9},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "E"},
				{"b", "A"},
				{"c", "F"},
				{"d", "G"},
				{"e", "D"},
				{"f", "C"},
				{"g", "B"},
				{"h", "H"},
				{"i", "I"},
			},
		},
		{
			Question: model.Question{
				ProblemDescription: `
ある会社では、給与は社員のレベルによる基本給、手当て（役職手当、単身赴任手当て）、勤務票からの情報（残業時間や欠勤など）で計算されます。役職手当て、単身赴任手当てはそれぞれ同時に2つ以上つくことはありません。
出社および退社の時間はタイムカードで管理されます。有給休暇取得や残業をするときは申請書が必要になります。`,
				Question:          "次のクラス図のa~jに該当するものを選択し、モデルを完成させなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram2.png?alt=media&token=7555a4dd-7e75-4d11-ae95-47a37aecbc53",
				ClassDiagramPlantUML: `@startuml
class a { }
class b { }
class c { }
class d {
日付
}
class e {
役職
}
class f { }
class g {
レベル
評価
}
class h {
出社時間
退社時間
}
class i {
日数
}
class j {
時間
}
a "0..2" --o "0.." b
a <|-- e
a <|-- f
b "0..1" --o "1" g
b "0..1" -- "1" c
c "1" o-- "0..31" h
c "1" o-- "0.." d
d <|-- i
d <|-- j
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "勤務票"},
				{ChoiceCode: "B", ChoiceText: "手当て"},
				{ChoiceCode: "C", ChoiceText: "給与明細書"},
				{ChoiceCode: "D", ChoiceText: "申請書"},
				{ChoiceCode: "E", ChoiceText: "有給休暇申請書"},
				{ChoiceCode: "F", ChoiceText: "社員"},
				{ChoiceCode: "G", ChoiceText: "残業申請書"},
				{ChoiceCode: "H", ChoiceText: "役職手当て"},
				{ChoiceCode: "I", ChoiceText: "単身赴任手当て"},
				{ChoiceCode: "J", ChoiceText: "タイムカード"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
				{LabelCode: "g", OrderId: 7},
				{LabelCode: "h", OrderId: 8},
				{LabelCode: "i", OrderId: 9},
				{LabelCode: "j", OrderId: 10},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "B"},
				{"b", "C"},
				{"c", "A"},
				{"d", "D"},
				{"e", "H"},
				{"f", "I"},
				{"g", "F"},
				{"h", "J"},
				{"i", "E"},
				{"j", "G"},
			},
		},
	}

	for _, s := range seeds {
		if err := createQuestion(db, s); err != nil {
			return err
		}
	}
	return nil
}

// 共通処理
func createQuestion(db *gorm.DB, seed QuestionSeed) error {
	q := seed.Question
	if err := db.Where("question = ?", q.Question).FirstOrCreate(&q).Error; err != nil {
		return err
	}

	// Choices
	for _, c := range seed.Choices {
		c.QuestionID = q.ID
		db.Where("choice_code = ? AND question_id = ?", c.ChoiceCode, q.ID).FirstOrCreate(&c)
	}

	// Labels
	for _, l := range seed.Labels {
		l.QuestionID = q.ID
		db.Where("label_code = ? AND question_id = ?", l.LabelCode, q.ID).FirstOrCreate(&l)
	}

	// Mappings
	for _, m := range seed.Mappings {
		var choice model.Choice
		var label model.Label
		if err := db.Where("choice_code = ? AND question_id = ?", m.ChoiceCode, q.ID).First(&choice).Error; err != nil {
			continue
		}
		if err := db.Where("label_code = ? AND question_id = ?", m.LabelCode, q.ID).First(&label).Error; err != nil {
			continue
		}
		am := model.AnswerMapping{QuestionID: q.ID, ChoiceID: choice.ID, LabelID: label.ID}
		db.Where("question_id = ? AND choice_id = ? AND label_id = ?", q.ID, choice.ID, label.ID).FirstOrCreate(&am)
	}

	return nil
}
