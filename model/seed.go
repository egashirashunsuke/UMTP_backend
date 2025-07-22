package model

import (
	"gorm.io/gorm"
)

func SeedDate(db *gorm.DB) error {
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
			{ChoiceCode: "A", ChoiceText: "診断指標"},
			{ChoiceCode: "B", ChoiceText: "基本データ"},
			{ChoiceCode: "C", ChoiceText: "診断結果"},
			{ChoiceCode: "D", ChoiceText: "会社情報"},
			{ChoiceCode: "E", ChoiceText: "財務データ"},
			{ChoiceCode: "F", ChoiceText: "自社情報"},
			{ChoiceCode: "G", ChoiceText: "業種"},
		},
	}
	if err := db.
		Where("question = ?", question.Question).
		FirstOrCreate(&question).
		Error; err != nil {
		return err
	}
	return nil
}
