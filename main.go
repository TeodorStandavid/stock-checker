package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

var stockSymbol string
var nDays int
var apiKey string

func main() {

	var err error

	stockSymbol = os.Getenv("SYMBOL")
	apiKey = os.Getenv("API_KEY")
	nDays, err = strconv.Atoi(os.Getenv("NDAYS")) 

	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()

	r.GET("/", GetSymbolPrices)

	r.Run()

}

func GetSymbolPrices(c *gin.Context) {

	lastNClosePrices := lastNClosePrices(nDays)
	averageStockPrice := math.Round(averageStockPrice(lastNClosePrices)) 
	c.JSON(http.StatusOK, gin.H{"data": lastNClosePrices, "average": averageStockPrice})
	// c.JSON(http.StatusOK, gin.H{"average": averageStockPrice})
}

func lastNClosePrices(lastNDays int) []float64 {

	closingDatedAndPrices := getClosingPrices()

	closingPriceArray := []float64{}

	dates := []time.Time{}

	for date := range closingDatedAndPrices {
		dates = append(dates, parseDate(date))
	}

	sort.Slice(dates, func(i, j int) bool {
		return dates[i].After(dates[j])
	})

	for i := 0; i < lastNDays; i++ {

		date := dates[i].Format("2006-01-02")

		closingPriceArray = append(closingPriceArray, closingDatedAndPrices[date])

	}

	return closingPriceArray

}

func averageStockPrice(prices []float64) float64 {

	sum := float64(0)

	for i := 0; i < len(prices); i++ {
		sum += prices[i]
	}
	return float64(sum) / float64(len(prices))
}

func getClosingPrices() map[string]float64 {

	client := &http.Client{}

	requestURL := "https://www.alphavantage.co/query?apikey=" + apiKey + "&function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + stockSymbol + "&datatype=json"

	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject interface{}
	json.Unmarshal(bodyBytes, &responseObject)
	AssertedResponseObject := responseObject.(map[string]interface{})["Time Series (Daily)"].(map[string]interface{})

	closingPriceArray := make(map[string]float64)

	for k, v := range AssertedResponseObject {

		v = v.(map[string]interface{})["4. close"].(string)

		v, err := strconv.ParseFloat(v.(string), 64)

		if err != nil {
			fmt.Println(err)
		}

		closingPriceArray[k] = v

	}

	return closingPriceArray

}

func parseDate(date string) time.Time {
	layout := "2006-01-02"
	dateParsed, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return dateParsed
}
