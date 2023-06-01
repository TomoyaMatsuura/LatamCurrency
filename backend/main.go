package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type pageInfo struct {
	Currency   string    `json:"currency"`
	FY         string    `json:"FY"`
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
func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}

func main() {
	currency := [8]string{"JPY", "BRL", "MXN", "ARS", "CLP", "PEN", "COP", "BOB"}

	// 対象8通貨をループ
	for i := 0; i < len(currency); i++ {

		// 今年の為替
		year, month, _ := time.Now().Date()
		startOfYear := time.Date(year, time.April, 1, 0, 0, 0, 0, time.UTC)
		unixTime1 := startOfYear.Unix()
		unixTime1Str := strconv.Itoa(int(unixTime1))

		thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		unixTime2 := thisMonth.Unix()
		unixTime2Str := strconv.Itoa(int(unixTime2))

		url := "https://finance.yahoo.com/quote/USD" + currency[i] + "%3DX/history?period1=" + unixTime1Str + "&period2=" + unixTime2Str + "&interval=1mo&filter=history&frequency=1mo&includeAdjustedClose=true"

		// 前年度の為替
		lastFiscalYearStart := time.Date(year-1, time.April, 1, 0, 0, 0, 0, time.UTC)
		lastFiscalYearEnd := time.Date(year, time.March, 31, 0, 0, 0, 0, time.UTC)
		unixTime1OfLastYear := lastFiscalYearStart.Unix()
		unixTime1StrOfLastYear := strconv.Itoa(int(unixTime1OfLastYear))
		unixTime2OfLastYear := lastFiscalYearEnd.Unix()
		unixTime2StrOfLastYear := strconv.Itoa(int(unixTime2OfLastYear))

		url2 := "https://finance.yahoo.com/quote/USD" + currency[i] + "%3DX/history?period1=" + unixTime1StrOfLastYear + "&period2=" + unixTime2StrOfLastYear + "&interval=1mo&filter=history&frequency=1mo&includeAdjustedClose=true"

		p := &pageInfo{
			Currency: currency[i],
			FY:       fmt.Sprintf("FY%d", year),
		}

		p2 := &pageInfo{
			Currency: currency[i],
			FY:       fmt.Sprintf("FY%d", year-1),
		}

		// コントローラを作成
		c := colly.NewCollector()
		c2 := colly.NewCollector()

		// タイトルを取得
		c.OnHTML("title", func(e *colly.HTMLElement) {
			p.Title = e.Text
			fmt.Println(e.Text)
		})

		c2.OnHTML("title", func(e *colly.HTMLElement) {
			p2.Title = e.Text
			fmt.Println(e.Text)
		})

		// 為替を取得
		c.OnHTML("table[data-test='historical-prices']", func(e *colly.HTMLElement) {
			e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
				date := el.ChildText("td:nth-child(1)")
				rate := el.ChildText("td:nth-child(5)")

				month := strings.Split(date, " ")[0]
				p.Month = append(p.Month, month)

				// レートをfloat64として解析し、小数点以下2桁に丸める
				r, err := strconv.ParseFloat(strings.ReplaceAll(rate, ",", ""), 64)
				if err != nil {
					log.Println("Error parsing rate:", err)
					return
				}
				r = roundToTwoDecimalPlaces(r)
				p.Rate = append(p.Rate, r)
			})
		})

		c2.OnHTML("table[data-test='historical-prices']", func(e *colly.HTMLElement) {
			e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
				date := el.ChildText("td:nth-child(1)")
				rate := el.ChildText("td:nth-child(5)")

				month := strings.Split(date, " ")[0]
				p2.Month = append(p2.Month, month)

				// レートをfloat64として解析し、小数点以下2桁に丸める
				r, err := strconv.ParseFloat(strings.ReplaceAll(rate, ",", ""), 64)
				if err != nil {
					log.Println("Error parsing rate:", err)
					return
				}
				r = roundToTwoDecimalPlaces(r)
				p2.Rate = append(p2.Rate, r)
			})
		})

		// URLを出力
		c.OnRequest(func(r *colly.Request) {
			p.URL = r.URL.String()
			fmt.Println("Visiting URL:", r.URL.String())
		})

		c2.OnRequest(func(r *colly.Request) {
			p2.URL = r.URL.String()
			fmt.Println("Visiting URL:", r.URL.String())
		})

		// Responseのステータスコードを取得
		c.OnResponse(func(r *colly.Response) {
			p.StatusCode = r.StatusCode
			fmt.Println("StatusCode:", r.StatusCode)
		})

		c2.OnResponse(func(r *colly.Response) {
			p2.StatusCode = r.StatusCode
			fmt.Println("StatusCode:", r.StatusCode)
		})

		// Response時にエラーがあった場合同様に出力
		c.OnError(func(r *colly.Response, err error) {
			p.StatusCode = r.StatusCode
			log.Println("error:", r.StatusCode, err)
		})

		c2.OnError(func(r *colly.Response, err error) {
			p2.StatusCode = r.StatusCode
			log.Println("error:", r.StatusCode, err)
		})

		c.Visit(url)
		c2.Visit(url2)

		c.Wait()
		c2.Wait()

		savePageJson(fmt.Sprintf("rate_%s.json", currency[i]), p)
		savePageJson(fmt.Sprintf("rateLy_%s.json", currency[i]), p2)
	}
}
