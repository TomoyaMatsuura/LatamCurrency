package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type pageInfo struct {
	Currency   string    `json:"currency"`
	StatusCode int       `json:"statusCode"`
	URL        string    `json:"url"`
	Title      string    `json:"title"`
	Month      []string  `json:"month"`
	Rate       []float64 `json:"rate"`
}

// JSONファイル作成用の関数
func savePageJson(fName string, p *pageInfo) {
	// JSONファイルの作成
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// エンコード
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	// JSON形式に変換
	b, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(b))
}

// 小数点下2桁に丸める関数
func roundToFourDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}

func main() {
	currency := [8]string{"JPY", "BRL", "MXN", "ARS", "CLP", "PEN", "COP", "BOB"}

	// 対象8通貨をループ
	for i := 0; i < len(currency); i++ {
		url := "https://finance.yahoo.com/quote/USD" + currency[i] + "%3DX/history?interval=1mo&filter=history&frequency=1mo"

		p := &pageInfo{
			Currency: currency[i],
		}

		// コントローラを作成
		c := colly.NewCollector()

		// タイトルを取得
		c.OnHTML("title", func(e *colly.HTMLElement) {
			p.Title = e.Text
			fmt.Println(e.Text)
		})

		// 為替を取得
		c.OnHTML("table[data-test='historical-prices']", func(e *colly.HTMLElement) {
			e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
				date := el.ChildText("td:nth-child(1)")
				rate := el.ChildText("td:nth-child(5)")

				month := strings.Split(date, " ")[0]
				p.Month = append(p.Month, month)

				// レートをfloat64として解析し、小数点以下4桁に丸める
				r, err := strconv.ParseFloat(strings.ReplaceAll(rate, ",", ""), 64)
				if err != nil {
					log.Println("Error parsing rate:", err)
					return
				}
				r = roundToFourDecimalPlaces(r)
				p.Rate = append(p.Rate, r)
			})
		})

		// URLを出力
		c.OnRequest(func(r *colly.Request) {
			p.URL = r.URL.String()
			fmt.Println("Visiting URL:", r.URL.String())
		})

		// Responseのステータスコードを取得
		c.OnResponse(func(r *colly.Response) {
			p.StatusCode = r.StatusCode
			fmt.Println("StatusCode:", r.StatusCode)
		})

		// Response時にエラーがあった場合同様に出力
		c.OnError(func(r *colly.Response, err error) {
			p.StatusCode = r.StatusCode
			log.Println("error:", r.StatusCode, err)
		})

		c.Visit(url)
		c.Wait()

		savePageJson(fmt.Sprintf("rate_%s.json", currency[i]), p)
	}
}
