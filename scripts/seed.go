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
				OrderIndex: 1,
				ProblemDescription: `
市営駐輪場の利用申し込みシステムです。
この市には第1と第2の駐輪場があります。それぞれの駐輪場で、月額料金は異なります。
市営駐輪場の利用を申し込むには、事前に会員登録する必要があります。会員登録では、氏名、住所、電話番号、メールアドレスを登録します。会員登録をすると、登録番号とパスワードがメールで送付されます。駐輪場の申し込み画面を開き、登録番号とパスワードを入力し、ログインしたあとに、希望の駐輪場を選択し利用申請をします。希望者が多数の場合は、抽選が行われます。抽選が外れると、次回の抽選で当選しやすくなります。`,
				Question:          "次のクラス図のa〜gに該当するものを選択し、モデルを完成させなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/classdiagram.png?alt=media&token=76ecd79f-1317-42a5-bcc9-c31278982abc",
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
				OrderIndex: 2,
				ProblemDescription: `
経営診断システムの開発を検討しています。
診断する会社についての以下の情報を入力することで、会社の経営診断結果が表示されます。
・基本データ（業種、従業員数など）
・財務データ（流動資産、固定資産、売上、利益など）
診断は、業種別に蓄積されている既存の蓄積データと比較することにより行われます。今回診断のため入力した会社データも、この蓄積データに追加されます。診断結果は、総合評価と、収益性、効率性、安全性、成長性のそれぞれの診断指標で表示されます。診断は、1社1回のみ行うことができます。`,
				Question:          "次のクラス図のa~gに該当するものを選択しなさい。なお、該当する選択肢が複数ある場合は、選択肢のアルファベットの早い順に選択しなさい。",
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
				OrderIndex: 3,
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
class d {
}
class e {
}
class g {
}
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
				OrderIndex: 4,
				ProblemDescription: `
ある中学校です。
1年〜3年まで3学年あり、1学年にはそれぞれ5つのクラスがあります。
1つのクラスは、担任の先生1人と副担任の先生1人が担当します。
担任の先生は1つのクラスのみ担当しますが、副担任の先生は、2つのクラスを担当することもあります。また、先生には担任も副担任も担当しない先生もいます。
先生は必ず、専門教科を1つ持っています。1教科は必ず1人以上の先生が担当しています。
また、先生は、1つのクラブの顧問を担当することもあれば担当者しないこともあります。クラブは必ず1人の先生が顧問をしています。`,
				Question:          "次のクラス図のa~eに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question7.png?alt=media&token=d8775145-ceea-4641-a8e5-b4f391518ae6",
				ClassDiagramPlantUML: `@startuml
class 学校 {
}
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
学校 “1” o-- "3" a
a “1” o-- "5" b 
b "0..1 f" -- "1 h" c 
b "0..2 g" -- "1 i" c
c "1" -- "0..1 j" d 
c "1..*" -- "1 k" e 
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "クラブ"},
				{ChoiceCode: "B", ChoiceText: "先生"},
				{ChoiceCode: "C", ChoiceText: "クラス"},
				{ChoiceCode: "D", ChoiceText: "学年"},
				{ChoiceCode: "E", ChoiceText: "教科"},
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
				{"c", "B"},
				{"d", "A"},
				{"e", "E"},
			},
		},
		{
			Question: model.Question{
				OrderIndex: 5,
				ProblemDescription: `
鉄道会社は、それぞれ複数の車両と路線を持っています。路線は複数の駅で構成されています。
鉄道会社が同じならば、駅は異なる路線で共有されることもあります。
経路は、出発駅と到着駅から構成されます。出発駅と到着駅は同一路線の場合もありますが、異なる路線の場合もあります。`,
				Question:          "次のクラス図のa〜eに該当するものを選択し、モデルを完成させなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question9.png?alt=media&token=e25f8bc6-e4f3-4b80-bf39-83ecc400dbf4",
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
a "1" o-- "1..*" b
a "1" o-- "1..*" c
c "1..*" -- "1..*" d
c "1..*" o-- "1..*" e
d "1..*" -- "1 出発駅" e
d "1..*" -- "1 到着駅" e
e "0..1 Prev" -- "0..1 Next" e
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "駅"},
				{ChoiceCode: "B", ChoiceText: "鉄道会社"},
				{ChoiceCode: "C", ChoiceText: "路線"},
				{ChoiceCode: "D", ChoiceText: "経路"},
				{ChoiceCode: "E", ChoiceText: "車両"},
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
				{"a", "B"},
				{"b", "E"},
				{"c", "C"},
				{"d", "D"},
				{"e", "A"},
			},
		},
		{
			Question: model.Question{
				OrderIndex: 6,
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
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question10.png?alt=media&token=0e0e44a3-3209-43f1-afda-8e9149feda34",
				ClassDiagramPlantUML: `@startuml
class 生徒 {
氏名
住所
学校種別
}
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
				OrderIndex: 7,
				ProblemDescription: `
ある会社では、トレーニングを提供しています。コース開発者は、コースの企画、設計および教材の開発を行います。完成したコースは、日程を決めて受講者の募集をします。申込みが3人以上いれば実施されます。実施時は1人のメイン講師が行い、補助講師が1人付くことがあります。`,
				Question:          "次のクラス図のa〜fに該当するものを選択し、モデルを完成させなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question11.png?alt=media&token=cdfccf86-4bec-4136-b77d-dc42eb67e1c5",
				ClassDiagramPlantUML: `@startuml
class a {
}
class b {
}
class c {
前提知識
期間
}
class d {
日程
}
class e {
氏名
電話番号
住所
}
class f {
氏名
}
a "1..*" -- "1..*" c
c "0..*" -- "1..* コース開発者" f
c "1" -- "0..*" d
b "0..*" -- "1" d
b "1..*" -- "1 申込者" e
e "3..*" -- "0..*" d
d "0..*" -- "1 メイン" f 
d "0..*" -- "0..1 補助" f 
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "講師"},
				{ChoiceCode: "B", ChoiceText: "受講者"},
				{ChoiceCode: "C", ChoiceText: "実施コース"},
				{ChoiceCode: "D", ChoiceText: "申込み"},
				{ChoiceCode: "E", ChoiceText: "コース"},
				{ChoiceCode: "F", ChoiceText: "教材"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "F"},
				{"b", "D"},
				{"c", "E"},
				{"d", "C"},
				{"e", "B"},
				{"f", "A"},
			},
		},
		{
			Question: model.Question{
				OrderIndex: 8,
				ProblemDescription: `
小売店は、卸売業者からお酒を仕入れ、顧客に販売します。
卸売業者は、メーカーからお酒を仕入れます。卸売業者と小売店は酒類販売業免許を持っている必要があります。卸売業者は酒類卸売業免許が、小売店は酒類小売業免許が必要になります。
酒類卸売業免許・・・酒類販売業者に販売するために必要な免許
酒類小売業免許・・・酒類を小売店など販売するために必要な免許`,
				Question:          "次のクラス図のa〜hに該当するものを選択し、モデルを完成させなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question12.png?alt=media&token=3e2d0350-6dfd-46d9-b7b4-4a0c1e78ef85",
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
class 個人消費者 {
}
a <|-- d
a <|-- c
d "1" --o "0..*" e
b "1..* 製造元" -- "1..*" e
e "1..* 仕入元" -- "1..*" f
c "1" --o "0..*" f
f "0..*" -- "0..* 販売先" g
g <|-- 個人消費者
g <|-- h
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "卸売業者"},
				{ChoiceCode: "B", ChoiceText: "酒類販売業免許"},
				{ChoiceCode: "C", ChoiceText: "メーカー"},
				{ChoiceCode: "D", ChoiceText: "酒類小売業免許"},
				{ChoiceCode: "E", ChoiceText: "顧客"},
				{ChoiceCode: "F", ChoiceText: "酒類卸売業免許"},
				{ChoiceCode: "G", ChoiceText: "小売店"},
				{ChoiceCode: "H", ChoiceText: "酒場・料理店"},
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
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "B"},
				{"b", "C"},
				{"c", "D"},
				{"d", "F"},
				{"e", "A"},
				{"f", "G"},
				{"g", "E"},
				{"h", "H"},
			},
		},
		{
			Question: model.Question{
				OrderIndex: 9,
				ProblemDescription: `
感染症は、病原体が原因で発生します。病原体には、寄生虫、細菌、真菌、ウイルスがあります。
感染経路には、飛沫感染、空気感染、接触感染、経口感染があります。
治療法には、薬剤によるものと外科的治療があります。また治療法のほか、予防法を検討する必要があります。
すべての感染症に関して、感染経路と病原体は判明していますが、治療法と予防法については、ないものもあります。`,
				Question:          "次のクラス図のa〜gに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question13.png?alt=media&token=25118d90-06b2-43da-a5d9-2eb0ee17e2a1",
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
class 病原体 {
}
class 外科的治療 {
}
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
				OrderIndex: 10,
				ProblemDescription: `
ある自転車店のパンク修理の費用を考えます。
パンク修理の費用は工賃と部品代の合計になります。
パンク修理は虫ゴムの交換のみ、パンク修理用パッチ処理、チューブ交換、チューブとタイヤの交換があります。虫ゴムの交換はどのパンク修理でも必ず行います。また、チューブだけの交換とチューブとタイヤの交換では工賃は変わりませんが、前輪と後輪では工賃は変わります。
後輪の場合は、後輪のブレーキ調整が必要です。またギア付き自転車の場合、後輪ギア調整も必要です。
修理は、前輪か後輪のどちらか1か所とします。`,
				Question:          "次のクラス図のa〜iに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question14.png?alt=media&token=1eecef44-37fc-4459-a8db-231d6d81547f",
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
				OrderIndex: 11,
				ProblemDescription: `
電子的に地図を参照することができます。
GPSから現在の位置を取得し、地図上に表示することができます。
GPSから1秒に1回、緯度、経度、高度を取得します。
GPSからの位置が2つあれば、区間となります。経路は1つ以上の区間の連続となります。地図は、地図イメージとキャリブレーションデータ（地図イメージの座標点と緯度、経度を関係付けるデータ）から構成されます。
また、地図上の任意の場所を設定ポイントとして、登録することができます。`,
				Question:          "次のクラス図のa〜gに該当するものを選択しなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question15.png?alt=media&token=8ac48e02-733d-4032-ad2f-3b99caa5d5ca",
				ClassDiagramPlantUML: `@startuml
class a {
a
}
class b {
b
}
class c {
c
}
class d {
d
}
class e {
e
}
class f {
f
}
class g {
g
}
class 設定ポイント {
h
}
a "0..1" o-- "1" b
b "2..*" --o "0..1" d
c "1" o-- "0..1" d
d o-- e
b "2" -- "0..1" f
f "0..1 前" -- "0..1 後" f
f "1..*" o-- "1" g
b "1" --o "0..1" 設定ポイント
a "0..1" -- "0..1" g
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "座標点"},
				{ChoiceCode: "B", ChoiceText: "区間"},
				{ChoiceCode: "C", ChoiceText: "キャリブレーション"},
				{ChoiceCode: "D", ChoiceText: "経路"},
				{ChoiceCode: "E", ChoiceText: "地図イメージ"},
				{ChoiceCode: "F", ChoiceText: "現在地"},
				{ChoiceCode: "G", ChoiceText: "位置"},
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
				{"a", "F"},
				{"b", "G"},
				{"c", "E"},
				{"d", "C"},
				{"e", "A"},
				{"f", "B"},
				{"g", "D"},
			},
		},
		{
			Question: model.Question{
				OrderIndex: 12,
				ProblemDescription: `
ある会社では、給与は社員のレベルによる基本給、手当て（役職手当、単身赴任手当て）、勤務票からの情報（残業時間や欠勤など）で計算されます。役職手当て、単身赴任手当てはそれぞれ同時に2つ以上つくことはありません。
出社および退社の時間はタイムカードで管理されます。有給休暇取得や残業をするときは申請書が必要になります。`,
				Question:          "次のクラス図のa~jに該当するものを選択し、モデルを完成させなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question17.png?alt=media&token=fb6eb863-77b7-47ce-8538-16163bb194e0",
				ClassDiagramPlantUML: `@startuml
class a { 
}
class b { 
}
class c { 
}
class d {
日付
}
class e {
役職
}
class f { 
}
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
		{
			Question: model.Question{
				OrderIndex: 13,
				ProblemDescription: `
ある会社の、従業員の通勤経路およびその費用管理のシステムです。
定期券がある場合は、6ヶ月の定期券代、定期券がない場合は片道運賃×往復×日数で計算します。
複数の交通機関がある場合は、その合計が通勤経路の合計金額になります。`,
				Question:          "次のクラス図のa~fに該当するものを選択し、モデルを完成させなさい。",
				ClassDiagramImage: "https://firebasestorage.googleapis.com/v0/b/umtp-learning.firebasestorage.app/o/question18.png?alt=media&token=2a1c82ec-cb7c-40b3-94b0-37adec4b18aa",
				ClassDiagramPlantUML: `@startuml
class a {
}
class b {
}
class c {
}
class d {
}
class JR {
}
class 私鉄 {
}
class バス {
}
a "1" o-- "1..*" c
b "1" --o "1" c
c “0..1 e” -- “0..1 f” c
c “0..*” -- “1” d
JR --|> d
私鉄 --|> d
バス --|> d
@enduml`,
			},
			Choices: []model.Choice{
				{ChoiceCode: "A", ChoiceText: "料金"},
				{ChoiceCode: "B", ChoiceText: "交通機関"},
				{ChoiceCode: "C", ChoiceText: "交通機関別経路"},
				{ChoiceCode: "D", ChoiceText: "通勤経路"},
				{ChoiceCode: "E", ChoiceText: "次"},
				{ChoiceCode: "F", ChoiceText: "前"},
			},
			Labels: []model.Label{
				{LabelCode: "a", OrderId: 1},
				{LabelCode: "b", OrderId: 2},
				{LabelCode: "c", OrderId: 3},
				{LabelCode: "d", OrderId: 4},
				{LabelCode: "e", OrderId: 5},
				{LabelCode: "f", OrderId: 6},
			},
			Mappings: []struct {
				LabelCode  string
				ChoiceCode string
			}{
				{"a", "D"},
				{"b", "A"},
				{"c", "C"},
				{"d", "B"},
				{"e", "E"},
				{"f", "F"},
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

func createQuestion(db *gorm.DB, seed QuestionSeed) error {
	q := seed.Question
	if err := db.Where("question = ?", q.ProblemDescription).FirstOrCreate(&q).Error; err != nil {
		return err
	}

	for _, c := range seed.Choices {
		c.QuestionID = q.ID
		db.Where("choice_code = ? AND question_id = ?", c.ChoiceCode, q.ID).FirstOrCreate(&c)
	}

	for _, l := range seed.Labels {
		l.QuestionID = q.ID
		db.Where("label_code = ? AND question_id = ?", l.LabelCode, q.ID).FirstOrCreate(&l)
	}

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
