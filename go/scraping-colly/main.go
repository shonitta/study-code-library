package main

import (
	"os"
	"fmt"
	"time"
	"encoding/csv"
	"path/filepath"

	"github.com/gocolly/colly/v2"
)

const (
    Domain = "www.bankers.co.jp"
    URL = "https://www.bankers.co.jp/investment/performance.html"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
	SaveDir = "outputs"
)

type TableRow struct {
	FundNo string
	Status string
	FundName string
	FundPageURL string
	PlannedInvestmentPeriod string
	ExpectedShare string
	AUM string
	RefundedPrincipal string
	CurrentAUM string
	ShareAmount string
	StartDate string
	EndDate string
	RefundType string
	ActualShare string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(Domain),
		colly.UserAgent(UserAgent),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob: Domain,
		Delay: 5 * time.Second,
		RandomDelay: 5 * time.Second,
	})

	table := make([]TableRow, 0)
    c.OnHTML("table.table-performance", func(e *colly.HTMLElement) {
		table = append(table, TableRow{
			"ファンドNo.",
			"ステータス",
			"ファンド名",
			"ファンドURL",
			"予定運用期間",
			"予定分配率（年率）",
			"運用額",
			"償還済み元本",
			"運用残高（償還繰越金含む）",
			"分配金額合計（税引前）",
			"運用開始日",
			"運用終了（予定）日",
			"元本償還/収益分配",
			"分配率実績（年率）",
		})
		e.ForEach("tbody > tr", func(_ int, el *colly.HTMLElement) {
			FundPageURL := el.ChildAttr("td:nth-of-type(3) > a", "href")
			table = append(table, TableRow{
				FundNo: el.ChildText("td:nth-of-type(1)"),
				Status: el.ChildText("td:nth-of-type(2)"),
				FundName: el.ChildText("td:nth-of-type(3)"),
				FundPageURL: FundPageURL,
				PlannedInvestmentPeriod: el.ChildText("td:nth-of-type(4)"),
				ExpectedShare: el.ChildText("td:nth-of-type(5)"),
				AUM: el.ChildText("td:nth-of-type(6)"),
				RefundedPrincipal: el.ChildText("td:nth-of-type(7)"),
				CurrentAUM: el.ChildText("td:nth-of-type(8)"),
				ShareAmount: el.ChildText("td:nth-of-type(9)"),
				StartDate: el.ChildText("td:nth-of-type(10)"),
				EndDate: el.ChildText("td:nth-of-type(11)"),
				RefundType: el.ChildText("td:nth-of-type(12)"),
				ActualShare: el.ChildText("td:nth-of-type(13)"),
			})
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	err := c.Visit(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll(SaveDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create(filepath.Join(SaveDir, "result.csv"))
	if err != nil {
		fmt.Println(err)
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
		}
		if err := w.Write(record); err != nil {
			fmt.Println(err)
			return
		}
	}
}
