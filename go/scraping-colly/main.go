package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const (
	Domain    = "www.bankers.co.jp"
	URL       = "https://" + Domain + "/investment/performance.html"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
	SaveDir   = "outputs"
)

type TableRow struct {
	FundNo                  string
	Status                  string
	FundName                string
	FundPageURL             string
	PlannedInvestmentPeriod string
	ExpectedShare           string
	AUM                     string
	RefundedPrincipal       string
	CurrentAUM              string
	ShareAmount             string
	StartDate               string
	EndDate                 string
	RefundType              string
	ActualShare             string
	FundFeatureText         string
	FundTermsConditionsText string
	FundScheduleText        string
	CollateralGuaranteeText string
	FundPurposeText         string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(Domain),
		colly.UserAgent(UserAgent),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  Domain,
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	table := make([]TableRow, 0)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})
	c.OnHTML("table.table-performance", func(e *colly.HTMLElement) {
		table = append(table, TableRow{
			"一覧表：ファンドNo.",
			"一覧表：ステータス",
			"一覧表：ファンド名",
			"一覧表：ファンドURL",
			"一覧表：予定運用期間",
			"一覧表：予定分配率（年率）",
			"一覧表：運用額",
			"一覧表：償還済み元本",
			"一覧表：運用残高（償還繰越金含む）",
			"一覧表：分配金額合計（税引前）",
			"一覧表：運用開始日",
			"一覧表：運用終了（予定）日",
			"一覧表：元本償還/収益分配",
			"一覧表：分配率実績（年率）",
			"個別ページ・ファンド概要：本ファンドのポイント",
			"個別ページ・ファンド概要：諸条件",
			"個別ページ・ファンド概要：スケジュール",
			"個別ページ・プロジェクト概要１：担保・保証",
			"個別ページ・プロジェクト概要１：借手資金使途",
		})

		e.ForEach("tbody > tr", func(_ int, el *colly.HTMLElement) {
			FundPageURL := el.ChildAttr("td:nth-of-type(3) > a", "href")

			// // TODO: デバッグ用なので削除する
			// // 特定のファンドページの挙動確認用
			// cond1 := FundPageURL != "https://www.bankers.co.jp/investment/lottery_application_entry.html?command=new&fund_id=516"
			// cond2 := FundPageURL != "https://www.bankers.co.jp/investment/investment_entry.html?command=new&fund_id=510"
			// cond3 := FundPageURL != "https://www.bankers.co.jp/investment/lottery_application_entry.html?command=new&fund_id=509"
			// if cond1 && cond2 && cond3 {
			// 	return
			// }

			row := TableRow{
				FundNo:                  el.ChildText("td:nth-of-type(1)"),
				Status:                  el.ChildText("td:nth-of-type(2)"),
				FundName:                el.ChildText("td:nth-of-type(3)"),
				FundPageURL:             FundPageURL,
				PlannedInvestmentPeriod: el.ChildText("td:nth-of-type(4)"),
				ExpectedShare:           el.ChildText("td:nth-of-type(5)"),
				AUM:                     el.ChildText("td:nth-of-type(6)"),
				RefundedPrincipal:       el.ChildText("td:nth-of-type(7)"),
				CurrentAUM:              el.ChildText("td:nth-of-type(8)"),
				ShareAmount:             el.ChildText("td:nth-of-type(9)"),
				StartDate:               el.ChildText("td:nth-of-type(10)"),
				EndDate:                 el.ChildText("td:nth-of-type(11)"),
				RefundType:              el.ChildText("td:nth-of-type(12)"),
				ActualShare:             el.ChildText("td:nth-of-type(13)"),
			}
			c2 := c.Clone()
			c2.OnRequest(func(r *colly.Request) {
				log.Println("Visiting", r.URL.String())
			})
			// 入れ子でファンドページにアクセスしないようにカウンターを設定
			counter2 := 0
			c2.OnHTML("div.main-contents__inner", func(eTabs *colly.HTMLElement) {
				if counter2 >= 1 {
					return
				}

				// 本ファンドのポイント & 諸条件 & スケジュール
				targetFlagFundFeatureText := false
				FundFeatureText := ""
				targetFlagFundTermsConditions := false
				FundTermsConditionsText := ""
				targetFlagFundSchedule := false
				FundScheduleText := ""
				eTabs.ForEach("section#tab1 > *", func(_ int, elTab1 *colly.HTMLElement) {
					if elTab1.Name == "h3" && elTab1.Text == "本ファンドのポイント" {
						targetFlagFundFeatureText = true
						log.Println("Found 本ファンドのポイント")
					} else if elTab1.Name == "h3" && targetFlagFundFeatureText {
						targetFlagFundFeatureText = false
					}
					if targetFlagFundFeatureText && elTab1.Name != "h3" {
						if elTab1.Name == "h4" && FundFeatureText != "" {
							FundFeatureText += "\n\n"
						}
						FundFeatureText += elTab1.Text
						if elTab1.Name == "h4" {
							FundFeatureText += "\n"
						}
					}

					if elTab1.Name == "h3" && elTab1.Text == "諸条件" {
						targetFlagFundTermsConditions = true
						log.Println("Found 諸条件")
					} else if elTab1.Name == "h3" && targetFlagFundTermsConditions {
						targetFlagFundTermsConditions = false
					}
					if targetFlagFundTermsConditions && elTab1.Name == "table" {
						elTab1.ForEach("tbody > tr", func(_ int, elTab1Tbl *colly.HTMLElement) {
							FundTermsConditionsText += strings.ReplaceAll(elTab1Tbl.ChildText("th"), "\n", "")
							FundTermsConditionsText += "\n"
							FundTermsConditionsText += elTab1Tbl.ChildText("td")
							FundTermsConditionsText += "\n-----\n"
						})
					}

					if elTab1.Name == "h3" && elTab1.Text == "スケジュール" {
						targetFlagFundSchedule = true
						log.Println("Found スケジュール")
					} else if elTab1.Name == "h3" && targetFlagFundSchedule {
						targetFlagFundSchedule = false
					}
					if targetFlagFundSchedule && elTab1.Name == "table" {
						elTab1.ForEach("tbody > tr", func(_ int, elTab1Tbl *colly.HTMLElement) {
							FundScheduleText += strings.ReplaceAll(elTab1Tbl.ChildText("th"), "\n", "")
							FundScheduleText += "\n"
							FundScheduleText += elTab1Tbl.ChildText("td")
							FundScheduleText += "\n-----\n"
						})
					}
				})

				// 担保・保証 & 借手資金使途
				targetFlagCollateralGuarantee := false
				CollateralGuaranteeText := ""
				targetFlagFundPurpose := false
				FundPurposeText := ""
				eTabs.ForEach("section#tab2 > *", func(_ int, elTab2 *colly.HTMLElement) {
					if elTab2.Name == "h3" && elTab2.Text == "担保・保証" {
						targetFlagCollateralGuarantee = true
						log.Println("Found 担保・保証")
					} else if elTab2.Name == "h3" && targetFlagCollateralGuarantee {
						targetFlagCollateralGuarantee = false
					}
					if targetFlagCollateralGuarantee && elTab2.Name != "h3" {
						if elTab2.Name == "h4" && CollateralGuaranteeText != "" {
							CollateralGuaranteeText += "\n\n"
						}
						CollateralGuaranteeText += elTab2.Text
						if elTab2.Name == "h4" {
							CollateralGuaranteeText += "\n"
						}
					}

					if elTab2.Name == "h3" && elTab2.Text == "借手資金使途" {
						targetFlagFundPurpose = true
						log.Println("Found 借手資金使途")
					} else if elTab2.Name == "h3" && targetFlagFundPurpose {
						targetFlagFundPurpose = false
					}
					if targetFlagFundPurpose && elTab2.Name != "h3" {
						if elTab2.Name == "h4" && FundPurposeText != "" {
							FundPurposeText += "\n\n"
						}
						FundPurposeText += elTab2.Text
						if elTab2.Name == "h4" {
							FundPurposeText += "\n"
						}
					}
				})

				row.FundFeatureText = FundFeatureText
				row.FundTermsConditionsText = FundTermsConditionsText
				row.FundScheduleText = FundScheduleText
				row.CollateralGuaranteeText = CollateralGuaranteeText
				row.FundPurposeText = FundPurposeText
				table = append(table, row)
				counter2++
			})
			err := c2.Visit(FundPageURL)
			if err != nil {
				log.Fatal(err)
				return
			}
		})
	})

	err := c.Visit(URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = os.MkdirAll(SaveDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	f, err := os.Create(filepath.Join(SaveDir, "result.csv"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	for _, row := range table {
		record := []string{
			row.FundNo,
			row.Status,
			row.FundName,
			row.FundPageURL,
			row.PlannedInvestmentPeriod,
			row.ExpectedShare,
			row.AUM,
			row.RefundedPrincipal,
			row.CurrentAUM,
			row.ShareAmount,
			row.StartDate,
			row.EndDate,
			row.RefundType,
			row.ActualShare,
			row.FundFeatureText,
			row.FundTermsConditionsText,
			row.FundScheduleText,
			row.CollateralGuaranteeText,
			row.FundPurposeText,
		}
		if err := w.Write(record); err != nil {
			log.Fatal(err)
			return
		}
	}
}
