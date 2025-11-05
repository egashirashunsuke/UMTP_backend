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
				Question: "次のクラス図のa〜gに該当するものを選択し、モデルを完成させなさい。",
				AnswerProcess: `
				{
"answers": [
{
"label": "a",
"choice": "G",
"choice_text": "会員リスト",
"reason": "aはbに対して集約(a o-- b)となっており、bを多数まとめる“入れ物”に相当する。問題文では会員登録で会員が管理されることが示唆され、会員を集約するクラスとして最も自然なのは「会員リスト」である。"
},
{
"label": "b",
"choice": "D",
"choice_text": "会員",
"reason": "会員は氏名・住所・電話・メールを登録し、登録番号とパスワードでログインして申請する主体。bはc(申請)と関連し、a(会員リスト)に集約されているため、申請の当事者である「会員」が妥当。"
},
{
"label": "c",
"choice": "E",
"choice_text": "申請",
"reason": "問題文では「ログイン後に希望の駐輪場を選択し利用申請をする」とある。cはb(会員)とe(駐輪場)の両方に関連しており、会員がどの駐輪場に対して行う行為・記録として「申請」に合致する。"
},
{
"label": "d",
"choice": "F",
"choice_text": "抽選結果",
"reason": "多数の場合に抽選が行われ、その結果が次回に影響するとある。c(申請)とdが関連しているため、各申請に対して結果が付随する「抽選結果」が適切。"
},
{
"label": "e",
"choice": "A",
"choice_text": "駐輪場",
"reason": "eにはf,gへの継承(e <|-- f, e <|-- g)があり、上位概念である必要がある。問題文では第1と第2の駐輪場が存在し、共通の上位概念は「駐輪場」。またc(申請)から対象施設への関連先にもなる。"
},
{
"label": "f",
"choice": "B",
"choice_text": "第1駐輪場",
"reason": "e(駐輪場)の下位クラスの一方であり、問題文の「第1駐輪場」に対応。月額料金が異なるという差異は下位クラスごとに表現される。"
},
{
"label": "g",
"choice": "C",
"choice_text": "第2駐輪場",
"reason": "e(駐輪場)のもう一方の下位クラスで、問題文の「第2駐輪場」に対応。第1と第2で料金が異なるため、継承で区別する設計に整合する。"
}
]
}
				`,
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
				Question: "次のクラス図のa~gに該当するものを選択しなさい。なお、該当する選択肢が複数ある場合は、選択肢のアルファベットの早い順に選択しなさい。",
				AnswerProcess: `
				{
  "answers": [
    {
      "label": "a",
      "choice": "G",
      "choice_text": "業種",
      "reason": "蓄積データ \"1\" o-- \"0..\" a から、蓄積データが多数の a を集約している。問題文の「業種別に蓄積されている既存の蓄積データ」と一致するため、a は業種。さらに a \"1\" o-- \"0..\" b で、業種の下に多数の会社情報が蓄積される構造とも整合。"
    },
    {
      "label": "b",
      "choice": "D",
      "choice_text": "会社情報",
      "reason": "a（業種） \"1\" o-- \"0..\" b より、各業種に多数ぶら下がる企業データの単位は b。加えて b \"1\" o-- c と b \"1\" o-- d で、b が基本データと財務データを部品として持つことから、b は会社情報が妥当。"
    },
    {
      "label": "c",
      "choice": "B",
      "choice_text": "基本データ",
      "reason": "b（会社情報）が c を集約しており（b \"1\" o-- c）、会社情報の構成要素として問題文にある「基本データ（業種、従業員数など）」に合致。また c \"j\" --o \"0..1\" g から、g（自社情報）が基本データを持つ（自社入力）関係とも一致。"
    },
    {
      "label": "d",
      "choice": "E",
      "choice_text": "財務データ",
      "reason": "b（会社情報）が d を集約（b \"1\" o-- d）し、会社情報の構成要素として問題文の「財務データ（流動資産、固定資産、売上、利益など）」に対応。さらに d \"k\" --o \"0..1\" g から、自社情報 g が財務データを持つ（入力データ）点とも一致。"
    },
    {
      "label": "e",
      "choice": "A",
      "choice_text": "診断指標",
      "reason": "e \"l\" --o \"1\" f で、f が e を複数集約している。問題文の「収益性、効率性、安全性、成長性のそれぞれの診断指標で表示」に対応し、診断結果が複数の診断指標を内包する構造に合致。"
    },
    {
      "label": "f",
      "choice": "C",
      "choice_text": "診断結果",
      "reason": "e を複数集約し（e --o f）、診断結果が指標の集合で表される点と一致。さらに f \"1\" --o \"1\" g で g と 1 対 1 の関係にあり、問題文の「診断は、1社1回のみ」に対応して、自社ごとに診断結果がちょうど1つであることを表す。"
    },
    {
      "label": "g",
      "choice": "F",
      "choice_text": "自社情報",
      "reason": "f（診断結果）と 1 対 1（f \"1\" --o \"1\" g）で「1社1回のみ」を満たす。さらに c（基本データ）および d（財務データ）を集約（c --o g, d --o g）しており、問題文の「今回診断のため入力した会社データ」に対応する“診断対象となる自社”の情報と解釈できる。"
    }
  ]
}
				`,
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
				Question: "次のクラス図のa~gに該当するものを選択しなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "B",
      "choice_text": "電子ビザ",
      "reason": "a には d と e が汎化の子として接続されており（a <|-- d, a <|-- e）、問題文の「電子ビザには、観光用と短期商用のものがあります」に対応する上位概念は電子ビザである。また b（申請）から a へ 0..1 の関連があり、申請の結果として電子ビザが発行されるという文脈にも合致する。"
    },
    {
      "label": "b",
      "choice": "C",
      "choice_text": "申請登録",
      "reason": "b には属性「参照番号」があり、問題文の「申請すると、参照番号が発行されます」に一致する。さらに b は a（電子ビザ）と 1 − 0..1 で関連しており、1 件の申請登録から最大 1 件の電子ビザが発行される関係を表している。"
    },
    {
      "label": "c",
      "choice": "D",
      "choice_text": "申請者",
      "reason": "c はクレジットカードおよび f（パスポート）を集約（c o-- f, c o-- クレジットカード）しており、問題文の「登録には、パスポートの情報およびクレジットカードの情報が必要」に対応して、それらの情報を持つ主体＝申請者が適切である。また b（申請登録）との関連が c 側 1、b 側 1..* となり、1 人の申請者が複数回申請でき、各申請は一意の申請者に紐づくという自然な関係になる。"
    },
    {
      "label": "d",
      "choice": "F",
      "choice_text": "観光用ビザ",
      "reason": "d は a（電子ビザ）のサブクラスであり、さらに d から g への 1 − 0..* の集約（d \"1\" o-- \"0..\" g）がある。問題文の「観光用のものは、有効期限内ならば何回でも使用できます」に対応して、複数回の使用履歴（g）を保持する必要があるのは観光用ビザであるため、d＝観光用ビザが妥当。"
    },
    {
      "label": "e",
      "choice": "G",
      "choice_text": "短期商用電子ビザ",
      "reason": "e は a（電子ビザ）のもう一方のサブクラスであり、観光用に対する対概念として短期商用電子ビザが該当する。d 側のみに履歴集約がある点は、観光用のみが複数回使用（履歴多数）を持つという仕様に整合するため、e＝短期商用電子ビザと判断できる。"
    },
    {
      "label": "f",
      "choice": "A",
      "choice_text": "パスポート",
      "reason": "f には属性「有効期限」があり、パスポートに有効期限がある点と一致する。さらに c（申請者）から f への集約（c o-- f）は、申請者がパスポート情報を持つという問題文の要件に合致する。クレジットカードは別クラスとして明示されているため、f はパスポートで確定できる。"
    },
    {
      "label": "g",
      "choice": "E",
      "choice_text": "履歴",
      "reason": "g は d（観光用ビザ）から 0..* で集約されている（d \"1\" o-- \"0..\" g）。問題文の「観光用のものは…何回でも使用できます」を設計に落とすと、使用の都度の記録＝履歴を多数持つ構造となるため、g＝履歴が最も整合的である。"
    }
  ]
}`,
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
				Question: "次のクラス図のa~eに該当するものを選択しなさい。",
				AnswerProcess: `
{
"answers": [
{
"label": "a",
"choice": "D",
"choice_text": "学年",
"reason": "学校 “1” o-- “3” a より、学校は3つのaを集約して持つ。問題文「1年〜3年まで3学年あり」と一致するため、a=学年。"
},
{
"label": "b",
"choice": "C",
"choice_text": "クラス",
"reason": "a “1” o-- “5” b より、1学年に5つのbが集約される。問題文「1学年にはそれぞれ5つのクラスがあります」と一致するため、b=クラス。"
},
{
"label": "c",
"choice": "B",
"choice_text": "先生",
"reason": "b “0..1 f” -- “1 h” c と b “0..2 g” -- “1 i” c の2本の関連があり、各クラス(b)にはそれぞれ1人のc（h=担任, i=副担任）が必須。一方c側は担任は0..1、副担任は0..2で、問題文「担任は1クラスのみ、副担任は最大2クラス、どちらも担当しない先生もいる」と一致。よってc=先生。"
},
{
"label": "d",
"choice": "A",
"choice_text": "クラブ",
"reason": "c “1” -- “0..1 j” d は、d（クラブ）側からは必ず1人のc（先生）が結び付き、c（先生）側からは0..1のd（クラブ）と結び付くことを示す。これは問題文「先生は1つのクラブの顧問を担当することもあればしないことも。クラブは必ず1人の先生が顧問」と一致。よってd=クラブ。"
},
{
"label": "e",
"choice": "E",
"choice_text": "教科",
"reason": "c “1..*” -- “1 k” e は、e（教科）には1人以上のc（先生）が担当し、c（先生）は必ず1つのe（教科）を持つことを示す。問題文「先生は必ず専門教科を1つ持つ。1教科は必ず1人以上の先生が担当」と一致。よってe=教科。"
}
]
}`,
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
				Question: "次のクラス図のa〜eに該当するものを選択し、モデルを完成させなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "B",
      "choice_text": "鉄道会社",
      "reason": "aはbとcに対して「1 o-- 1..*」の集約を持ち、1つのaが複数のb（車両）と複数のc（路線）を保持する。問題文の「鉄道会社は、それぞれ複数の車両と路線を持っています。」に一致するため、a=鉄道会社。"
    },
    {
      "label": "b",
      "choice": "E",
      "choice_text": "車両",
      "reason": "bはaから「1 o-- 1..*」で集約されており、1つの鉄道会社が複数の車両を持つ構造に合致する。他の関係を持たない点も、本文の範囲（会社が車両を保有）と整合するため、b=車両。"
    },
    {
      "label": "c",
      "choice": "C",
      "choice_text": "路線",
      "reason": "cはaから「1 o-- 1..*」で集約され（会社は複数の路線を持つ）、さらにe（駅）に対して「c 1..* o-- 1..* e」となり、路線が複数の駅で構成され、駅が複数路線に属し得る（共有可能）ことを表す。問題文の「路線は複数の駅で構成」「同じ鉄道会社なら駅は異なる路線で共有」に対応するため、c=路線。"
    },
    {
      "label": "d",
      "choice": "D",
      "choice_text": "経路",
      "reason": "dはe（駅）に対して「d 1..* -- 1 出発駅 e」「d 1..* -- 1 到着駅 e」となっており、各dが必ず1つの出発駅・1つの到着駅を持つ。これは「経路は、出発駅と到着駅から構成」に合致する。また「c 1..* -- 1..* d」により、経路が同一路線のみ（1本）にも異なる路線（複数）にも関わり得る点が「同一路線の場合も…異なる路線の場合も…」に対応するため、d=経路。"
    },
    {
      "label": "e",
      "choice": "A",
      "choice_text": "駅",
      "reason": "eはdとの関連端名が「出発駅」「到着駅」であり、駅の役割を明示している。またe同士の「0..1 Prev／0..1 Next」の自己関連は路線上での駅の前後関係（並び）を表し、c（路線）との集約「c 1..* o-- 1..* e」で路線が駅の集合であることを示す。これらは本文の駅の説明に合致するため、e=駅。"
    }
  ]
}`,
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
				Question: "クラス図のa〜eに該当するものを選択しなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "D",
      "choice_text": "希望家庭教師像",
      "reason": "生徒は希望（希望家庭教師像）を1〜3つ入力するとある。図ではa \"f\" --o \"1\" 生徒となっており、生徒側に集約（diamond）があり、aは必ず1人の生徒に属する（生徒側多重度1）。これは「生徒が複数の希望を保持する」関係に合致する。またaは次に示すエリアbを必ず1つ（a \"g\" o-- \"1\" b）持つため、希望に含まれる希望エリアの指定にも一致する。"
    },
    {
      "label": "b",
      "choice": "C",
      "choice_text": "エリア",
      "reason": "問題文の希望項目に「希望エリア（都道府県、市区町村群）」がある。図でa \"g\" o-- \"1\" bは、希望(a)がエリア(b)を必ず1つ含むことを表すため整合する。さらにb \"1\" o-- \"h\" dはエリア(b)が家庭教師(d)を集約しており、多重度\"1\"がb側に付くことで各家庭教師(d)が1つのエリアに属することを示す（エリアが家庭教師を束ねるという現実世界の構造に沿う）。"
    },
    {
      "label": "c",
      "choice": "E",
      "choice_text": "学校",
      "reason": "問題文では学校種別（小・中・高）ごとに選択可能な科目が異なるとある。図のc \"1\" o-- \"i\" eはcがeを集約し、c側多重度が1なので、各e（科目）は必ず1つのc（学校）に属することを意味する。これは「科目は特定の学校種別に属する」という要件に一致する。"
    },
    {
      "label": "d",
      "choice": "A",
      "choice_text": "家庭教師",
      "reason": "システムが提示する候補は家庭教師である。図でb \"1\" o-- \"h\" dは、エリア(b)が家庭教師(d)を集約して束ねる関係を示し、またd \"k\" -- \"j\" eにより家庭教師(d)が教えられる科目(e)との関連を持つ。これは家庭教師が担当可能科目を持つという要件と整合する。"
    },
    {
      "label": "e",
      "choice": "B",
      "choice_text": "科目",
      "reason": "問題文にある科目（小：国語/算数…、中：国語/数学…、高：国語系/数学系…）は学校種別ごとに異なる。図ではc \"1\" o-- \"i\" eにより科目(e)が学校(c)に属し、d \"k\" -- \"j\" eにより家庭教師(d)と科目(e)が関連付く。これは学校種別ごとに科目が定義され、家庭教師が教えられる科目を持つという構造を反映している。"
    }
  ]
}`,
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
				Question: "次のクラス図のa〜fに該当するものを選択し、モデルを完成させなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "F",
      "choice_text": "教材",
      "reason": "a と c（コース）が 1..* 対 1..* の関連で結ばれており、コースは複数の教材を用い、教材は複数コースで再利用され得る構造になっている。問題文でも「コース開発者は…教材の開発を行います」と教材の存在が明示されるため、a は教材が妥当。"
    },
    {
      "label": "b",
      "choice": "D",
      "choice_text": "申込み",
      "reason": "b は d（実施コース）に対して b側0..*—1 d と関連し、各申込みが1つの実施コースに紐づくことを示す。また e（受講者）に対して b側1..*—1 申込者 e となっており、各申込みはちょうど1人の申込者（受講者）によってなされ、1人の受講者は1件以上の申込みを持つ。問題文の「受講者の募集」「申込みが…」に一致するため、b は申込み。"
    },
    {
      "label": "c",
      "choice": "E",
      "choice_text": "コース",
      "reason": "c の属性が「前提知識」「期間」でありコース固有の情報に該当する。c は d（実施コース）と 1—0..* で結ばれ、1つのコースに複数の日程（実施）がある。さらに f（講師）と「1..* コース開発者」で結ばれ、各コースは1人以上の講師が開発するという問題文の「コース開発者…」に合致する。よって c はコース。"
    },
    {
      "label": "d",
      "choice": "C",
      "choice_text": "実施コース",
      "reason": "d の属性は「日程」であり、実際に開講される回（セッション）を表す。c（コース）と 1—0..* で結ばれ、1つのコースに複数の実施がある。f（講師）との関連では d側0..*—1 メイン と d側0..*—0..1 補助 となっており、「実施時は1人のメイン講師…補助講師が1人付くことがあります」に対応。e（受講者）との関連 e側3..*—0..* d は、各実施に受講者が3人以上必要という条件を表現している。従って d は実施コース。"
    },
    {
      "label": "e",
      "choice": "B",
      "choice_text": "受講者",
      "reason": "e の属性は氏名・電話番号・住所で、受講者の基本情報に相当。b（申込み）との関連で「1 申込者」とロール名が付いており、申込みの主体が受講者であることを示す。さらに d（実施コース）との e側3..*—0..* d により、各実施コースには3名以上の受講者が必要という問題文の要件を満たす。よって e は受講者。"
    },
    {
      "label": "f",
      "choice": "A",
      "choice_text": "講師",
      "reason": "f は c（コース）との関連で「コース開発者」というロールが付与され、講師がコース開発者であることを表す（各講師は0..*コースを開発、各コースは1..*講師が開発）。また d（実施コース）との関連では「メイン」1名必須・「補助」0..1名任意となっており、「実施時は1人のメイン講師…補助講師が1人付くことがあります」に一致する。従って f は講師。"
    }
  ]
}`,
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
				Question: "次のクラス図のa〜hに該当するものを選択し、モデルを完成させなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "B",
      "choice_text": "酒類販売業免許",
      "reason": "aはcとdの汎化元（a <|-- d, a <|-- c）で、問題文の「卸売業者と小売店は酒類販売業免許を持っている必要」が示す上位概念。c（小売）とd（卸売）の個別免許がa（酒類販売業免許）の下位に位置づくため。"
    },
    {
      "label": "b",
      "choice": "C",
      "choice_text": "メーカー",
      "reason": "bとeの関連にb側端名「製造元」が付与（b \"1..* 製造元\" -- \"1..*\" e）。卸売業者eにとっての製造元はメーカーであり、問題文の「卸売業者は、メーカーからお酒を仕入れます」と一致。多重度1..*同士も複数メーカーと複数卸売の関係に妥当。"
    },
    {
      "label": "c",
      "choice": "D",
      "choice_text": "酒類小売業免許",
      "reason": "aの下位（a <|-- c）で小売側の個別免許に該当。c \"1\" --o \"0..*\" fはf（小売店）側に集約（ダイアモンド）で「小売店は酒類小売業免許が必要」に合致。多重度は「各小売店が1つの小売免許を持つ（f→cが1）」ことを示す。"
    },
    {
      "label": "d",
      "choice": "F",
      "choice_text": "酒類卸売業免許",
      "reason": "aの下位（a <|-- d）で卸売側の個別免許に該当。d \"1\" --o \"0..*\" eはe（卸売業者）側に集約で「卸売業者は酒類卸売業免許が必要」に対応。多重度は「各卸売業者が1つの卸売免許を持つ（e→dが1）」ことを表す。"
    },
    {
      "label": "e",
      "choice": "A",
      "choice_text": "卸売業者",
      "reason": "b（メーカー）との関連でb端名が「製造元」、f（小売店）との関連でe端名が「仕入元」（e \"1..* 仕入元\" -- \"1..*\" f）。すなわちeは小売から見た仕入先であり、メーカーから仕入れる主体＝卸売業者。さらにd（酒類卸売業免許）を集約して保持する点も整合。"
    },
    {
      "label": "f",
      "choice": "G",
      "choice_text": "小売店",
      "reason": "e（卸売業者）との関連端名が「仕入元」で、fは仕入れる側＝小売店。g（顧客）との関連でg端名が「販売先」（f \"0..*\" -- \"0..* 販売先\" g）となっており、fが顧客へ販売する主体に該当。さらにc（酒類小売業免許）を集約して持つことも問題文と一致。"
    },
    {
      "label": "g",
      "choice": "E",
      "choice_text": "顧客",
      "reason": "gは「個人消費者」とhの汎化元（g <|-- 個人消費者, g <|-- h）で顧客の上位概念。fとの関連でg端名が「販売先」となっており、小売店の販売相手＝顧客を表す。"
    },
    {
      "label": "h",
      "choice": "H",
      "choice_text": "酒場・料理店",
      "reason": "g（顧客）の下位概念（g <|-- h）で、個人消費者以外の顧客カテゴリとして妥当。小売店の顧客には飲食店が含まれるため、酒場・料理店が該当。"
    }
  ]
}`,
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
				Question: "次のクラス図のa〜gに該当するものを選択しなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "C",
      "choice_text": "治療法",
      "reason": "a は d および「外科的治療」から汎化されている（a <|-- d, a <|-- 外科的治療）。問題文「治療法には、薬剤によるものと外科的治療があります」に合致し、a が上位概念の『治療法』である。さらに a は b（感染症）と関連し、治療法が感染症に紐づく点も問題文と一致。"
    },
    {
      "label": "b",
      "choice": "D",
      "choice_text": "感染症",
      "reason": "b は c（感染経路）への集約、病原体との関連、a（治療法）・f（予防法）との関連を持つ。問題文は『すべての感染症に関して、感染経路と病原体は判明』『治療法と予防法はないものもある』と述べ、中心となる対象が感染症であることから b は『感染症』。"
    },
    {
      "label": "c",
      "choice": "E",
      "choice_text": "感染経路",
      "reason": "b から c へ集約（b \"j\" o-- \"k\" c）があり、c は e に汎化されている（c <|-- e）。問題文に『感染経路には、飛沫感染、空気感染、接触感染、経口感染があります』とあり、c はその上位概念『感染経路』。"
    },
    {
      "label": "d",
      "choice": "F",
      "choice_text": "薬剤",
      "reason": "d は a（治療法）の下位（a <|-- d）。問題文『治療法には、薬剤によるものと外科的治療があります』より、治療法のサブタイプである d は『薬剤』。"
    },
    {
      "label": "e",
      "choice": "A",
      "choice_text": "飛沫感染 空気感染 接触感染 経口感染",
      "reason": "e は c（感染経路）の下位（c <|-- e）。問題文で列挙された4種の感染経路が c の具体的サブタイプであるため、e は『飛沫感染 空気感染 接触感染 経口感染』。"
    },
    {
      "label": "f",
      "choice": "G",
      "choice_text": "予防法",
      "reason": "b と f は関連（b \"l\" -- \"m\" f）しており、問題文『治療法のほか、予防法を検討する必要』『治療法と予防法については、ないものもあります』より、感染症に対して任意で結びつく概念は『予防法』。"
    },
    {
      "label": "g",
      "choice": "B",
      "choice_text": "寄生虫 細菌 真菌 ウイルス",
      "reason": "g は『病原体』の下位（病原体 <|-- g）。問題文『病原体には、寄生虫、細菌、真菌、ウイルスがあります』に対応し、g は病原体の具体的サブタイプ群『寄生虫 細菌 真菌 ウイルス』。"
    }
  ]
}`,
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
				Question: "次のクラス図のa〜iに該当するものを選択しなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "E",
      "choice_text": "パンク修理",
      "reason": "aはb(部品)を多数(n)集約し、かつc(工賃)を集約している（a \"1\" o-- \"n\" b と a \"1\" o-- \"j\" c）。問題文で「パンク修理の費用は工賃と部品代の合計」とあるため、工賃(c)と部品(b)を束ねるaは『パンク修理』に該当する。"
    },
    {
      "label": "b",
      "choice": "A",
      "choice_text": "部品",
      "reason": "bのサブクラスとして『タイヤ』『虫ゴム』『パンク修理用パッチ』『チューブ』がぶら下がる(b <|--)。いずれも部品であり、bはそれらを包括する上位概念『部品』となる。"
    },
    {
      "label": "c",
      "choice": "F",
      "choice_text": "パンク修理工賃",
      "reason": "a(パンク修理)がcを集約しており(a \"1\" o-- \"j\" c)、cが『虫ゴム交換』を必ず含む(c \"1\" o-- \"k\" 虫ゴム交換)。問題文の「虫ゴムの交換はどのパンク修理でも必ず行います」に対応する必須作業を内包することから、cは『パンク修理工賃』である。"
    },
    {
      "label": "d",
      "choice": "G",
      "choice_text": "パッチ処理",
      "reason": "c(工賃)がdを集約(c \"1\" o-- \"l\" d)。問題文の修理内容の一つ「パンク修理用パッチ処理」に対応する作業であり、部品としての『パンク修理用パッチ』（b配下）とは別に、作業としての『パッチ処理』が必要なためdはこれに該当。"
    },
    {
      "label": "e",
      "choice": "D",
      "choice_text": "車輪着脱",
      "reason": "c(工賃)がeを集約(c \"1\" o-- \"m\" e)し、eには前後輪での差異を表すサブクラスf, gがある(e <|-- f, e <|-- g)。問題文の「前輪と後輪では工賃は変わります」に対応する共通上位の作業が『車輪着脱』である。"
    },
    {
      "label": "f",
      "choice": "C",
      "choice_text": "前輪着脱",
      "reason": "e(車輪着脱)のサブクラスの一方が前輪側に該当する(e <|-- f)。前後で工賃が異なる差異を表すため、fは『前輪着脱』となる。"
    },
    {
      "label": "g",
      "choice": "B",
      "choice_text": "後輪着脱",
      "reason": "e(車輪着脱)のもう一方のサブクラスが後輪側(e <|-- g)。後輪の場合にのみ追加調整が必要であることを後続の関連で表しているため、gは『後輪着脱』。"
    },
    {
      "label": "h",
      "choice": "H",
      "choice_text": "後輪ギア調整",
      "reason": "g(後輪着脱)からhへの集約が0..1で任意(g \"0..1\" o-- \"0..1\" h)。問題文の「ギア付き自転車の場合、後輪ギア調整も必要」に一致する条件付き作業なので、hは『後輪ギア調整』。"
    },
    {
      "label": "i",
      "choice": "I",
      "choice_text": "後輪ブレーキ調整",
      "reason": "g(後輪着脱)からiへの集約が必須(g \"0..1\" o-- \"1\" i)となっており、問題文の「後輪の場合は、後輪のブレーキ調整が必要です」に対応する必須作業。ゆえにiは『後輪ブレーキ調整』。"
    }
  ]
}`,
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
				Question: "次のクラス図のa〜gに該当するものを選択しなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "F",
      "choice_text": "現在地",
      "reason": "aはb（位置）をちょうど1つ保持する集約 a \"0..1\" o-- \"1\" b であり、「GPSから現在の位置を取得し、地図上に表示する」に対応する“現在地”の概念が最も適合する。また a と g（経路）が相互に任意関連 a \"0..1\" -- \"0..1\" g で、現在地が経路に乗っている場合もあればそうでない場合もあることを表している。"
    },
    {
      "label": "b",
      "choice": "G",
      "choice_text": "位置",
      "reason": "問題文の「GPSから1秒に1回、緯度、経度、高度を取得」する“位置”に該当。b は f（区間）と「2 対 0..1」の関係 b \"2\" -- \"0..1\" f を持ち、区間はちょうど2つの位置からなるという記述「GPSからの位置が2つあれば、区間」に一致。また b は 設定ポイント と「1 対 0..1」の集約 b \"1\" --o \"0..1\" 設定ポイント を持ち、設定ポイントが1つの位置を登録するという要件に合致する。"
    },
    {
      "label": "c",
      "choice": "E",
      "choice_text": "地図イメージ",
      "reason": "「地図は、地図イメージとキャリブレーションデータから構成」に基づき、c は d（キャリブレーションデータ）を0..1だけ保持し、d は必ず1つの c に属する c \"1\" o-- \"0..1\" d。これは“ある地図イメージに対して0または1件のキャリブレーションデータが紐付く”という構造を表す。"
    },
    {
      "label": "d",
      "choice": "C",
      "choice_text": "キャリブレーションデータ",
      "reason": "キャリブレーションデータは「地図イメージの座標点と緯度・経度を関係付けるデータ」。d は e（座標点）を含む d o-- e とともに、b（位置）を2つ以上用いる b \"2..*\" --o \"0..1\" d で“画像上の座標点”と“地理的な位置”の対応を複数組持つことを表す。さらに d は c（地図イメージ）に属する c \"1\" o-- \"0..1\" d も整合している。"
    },
    {
      "label": "e",
      "choice": "A",
      "choice_text": "座標点",
      "reason": "d o-- e により、e は d（キャリブレーションデータ）の構成要素。問題文の「地図イメージの座標点と緯度、経度を関係付けるデータ」における“座標点”そのもの。"
    },
    {
      "label": "f",
      "choice": "B",
      "choice_text": "区間",
      "reason": "b（位置）と f の関連 b \"2\" -- \"0..1\" f が“区間は2つの位置からなる”に一致。さらに f の自己関連 f \"0..1 前\" -- \"0..1 後\" f が“区間の連続”を表す。後述の g（経路）との関係 f \"1..*\" o-- \"1\" g により、経路は1つ以上の区間から構成される要件にも合致。"
    },
    {
      "label": "g",
      "choice": "D",
      "choice_text": "経路",
      "reason": "f と g の関連 f \"1..*\" o-- \"1\" g は“経路は1つ以上の区間の連続”をそのまま表し、各区間は必ず1つの経路に属することを示す。また a（現在地）との任意関連 a \"0..1\" -- \"0..1\" g は、経路を辿っていない（紐付かない）場合があることも表している。"
    }
  ]
}`,
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
				Question: "次のクラス図のa~jに該当するものを選択し、モデルを完成させなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "B",
      "choice_text": "手当て",
      "reason": "aはe・fの共通上位で、b側に集約され0..2個まで関連（a \"0..2\" --o \"0..\" b）。問題文は「役職手当て、単身赴任手当ては…同時に2つ以上つくことはありません」とあり、同時に最大2種（各1つずつ）という多重度と一致するため、a=手当て。"
    },
    {
      "label": "b",
      "choice": "C",
      "choice_text": "給与明細書",
      "reason": "bはa（手当て）を最大2つ集約し、さらにc（勤務票）と1対1（b側1、c側0..1）で関連し、g（社員）に集約されている（b \"0..1\" --o \"1\" g）。給与明細は社員1人に対し当該期間で0..1通、勤務票1件から作成され、手当てを含むため、b=給与明細書。"
    },
    {
      "label": "c",
      "choice": "A",
      "choice_text": "勤務票",
      "reason": "cはh（タイムカード）を0..31件集約（c \"1\" o-- \"0..31\" h）し、d（申請書）も0..*集約（c \"1\" o-- \"0..\" d）。問題文の「勤務票からの情報（残業時間や欠勤）」「出社および退社の時間はタイムカード」「有給・残業は申請書が必要」と一致するため、c=勤務票。"
    },
    {
      "label": "d",
      "choice": "D",
      "choice_text": "申請書",
      "reason": "dは属性に日付を持ち、i（有給休暇申請書）、j（残業申請書）の上位クラス（d <|-- i, d <|-- j）。問題文「有給休暇取得や残業をするときは申請書が必要」に対応し、d=申請書。"
    },
    {
      "label": "e",
      "choice": "H",
      "choice_text": "役職手当て",
      "reason": "eはa（手当て）の下位で属性に「役職」を持つため、役職に紐づく手当てであることが明確。よってe=役職手当て。"
    },
    {
      "label": "f",
      "choice": "I",
      "choice_text": "単身赴任手当て",
      "reason": "fはa（手当て）の下位で、もう一方の特定手当てとして妥当。問題文にある2種のうち役職手当てがeであるため、f=単身赴任手当て。"
    },
    {
      "label": "g",
      "choice": "F",
      "choice_text": "社員",
      "reason": "gは属性に「レベル」「評価」を持ち、問題文「給与は社員のレベルによる基本給」に合致。b（給与明細書）がgに0..1で集約され（b \"0..1\" --o \"1\" g）、給与明細は社員に属するため、g=社員。"
    },
    {
      "label": "h",
      "choice": "J",
      "choice_text": "タイムカード",
      "reason": "hは属性に「出社時間」「退社時間」を持ち、問題文「出社および退社の時間はタイムカードで管理」に一致。c（勤務票）が月内0..31件のhを集約している点も妥当。よってh=タイムカード。"
    },
    {
      "label": "i",
      "choice": "E",
      "choice_text": "有給休暇申請書",
      "reason": "iはd（申請書）の下位で属性に「日数」を持つ。日数を扱う申請は有給休暇が自然で、問題文「有給休暇取得…は申請書が必要」に合致するため、i=有給休暇申請書。"
    },
    {
      "label": "j",
      "choice": "G",
      "choice_text": "残業申請書",
      "reason": "jはd（申請書）の下位で属性に「時間」を持つ。時間を扱う申請は残業申請が該当し、問題文「残業…は申請書が必要」に対応。よってj=残業申請書。"
    }
  ]
}`,
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
				Question: "次のクラス図のa~fに該当するものを選択し、モデルを完成させなさい。",
				AnswerProcess: `
{
  "answers": [
    {
      "label": "a",
      "choice": "D",
      "choice_text": "通勤経路",
      "reason": "aはcに対して1 o-- 1..*の集約であり、複数の交通機関の合計が通勤経路になるという記述と一致する。よって、複数のc（交通機関別経路）から構成される全体は通勤経路＝aである。"
    },
    {
      "label": "b",
      "choice": "A",
      "choice_text": "料金",
      "reason": "bはcに対して1 --o 1の集約（c側にダイヤモンド）で、各c（交通機関別経路）が1つの料金を持つ構造。問題文の計算規則（定期券の有無で計算し、複数交通機関なら合算）から、料金は経路単位（交通機関別経路単位）で保持・計算されるのが自然で、b＝料金が妥当。"
    },
    {
      "label": "c",
      "choice": "C",
      "choice_text": "交通機関別経路",
      "reason": "cはdに対して0..* -- 1の関連で、各cがちょうど1つのd（交通機関）に紐づく。さらにc同士が0..1同士で双方向連鎖（前/次）しており、通勤経路を構成する各区間（交通機関別の区間）を表すと解釈できるため、c＝交通機関別経路。"
    },
    {
      "label": "d",
      "choice": "B",
      "choice_text": "交通機関",
      "reason": "dをスーパークラスとしてJR・私鉄・バスが汎化でぶら下がっているため、dは交通機関の一般概念。問題文の「複数の交通機関がある場合…合計」記述とも整合。"
    },
    {
      "label": "e",
      "choice": "E",
      "choice_text": "次",
      "reason": "c同士の自己関連に0..1 e と 0..1 f の役割名があり、通勤経路内の順序付けを表す。各区間は高々1つの次区間を持つので、片側の役割名は「次」が適切。ここではeを「次」とする。"
    },
    {
      "label": "f",
      "choice": "F",
      "choice_text": "前",
      "reason": "自己関連のもう一方は、各区間が高々1つの前区間を持つことを表すための役割名であり、「前」が該当する。よってf＝前。"
    }
  ]
}`,
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
