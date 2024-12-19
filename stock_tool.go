package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

// NewsItem represents a single news entry with title and description
type NewsItem struct {
	Title   string // Title of the news
	Summary string // Description of the news
	Ticker  string
	Time    string
	Image   string
}

// Historical_data represents
type Historical_data struct {
	Date   string  // Title of the news
	Open   float64 // Description of the news
	Close  float64
	Low    float64
	High   float64
	Volume int
}

// Statistics represents all the statistical and financial data for a stock
type Statistics struct {
	// Financial Highlights
	FiscalYearEnd              string  // Fiscal year end
	MostRecentQuarter          string  // Most recent quarter
	ProfitMargin               float64 // Profit margin
	OperatingMargin            float64 // Operating margin (ttm)
	ReturnOnAssets             float64 // Return on assets (ttm)
	ReturnOnEquity             float64 // Return on equity (ttm)
	Revenue                    string  // Total revenue (ttm)
	RevenuePerShare            float64 // Revenue per share (ttm)
	QuarterlyRevenueGrowth     float64 // Quarterly revenue growth (yoy)
	EBITDA                     string  // EBITDA
	NetIncomeAvailableToCommon string  // Net income to common (ttm)
	DilutedEPS                 float64 // Diluted EPS (ttm)
	QuarterlyEarningsGrowth    string  // Quarterly earnings growth (yoy)
	TotalCash                  string  // Total cash (mrq)
	TotalCashPerShare          float64 // Total cash per share (mrq)
	TotalDebt                  string  // Total debt (mrq)
	TotalDebtEquity            float64 // Total debt/equity ratio (mrq)
	CurrentRatio               float64 // Current ratio (mrq)
	BookValuePerShare          float64 // Book value per share (mrq)
	OperatingCashFlow          string  // Operating cash flow (ttm)
	LeveredFreeCashFlow        string  // Levered free cash flow (ttm)

	// Trading Information
	Beta                       float64 // Beta (5Y monthly)
	FiftyTwoWeekRange          string  // 52-week range
	SP50052WeekChange          float64 // S&P 500 52-week change
	FiftyTwoWeekHigh           float64 // 52-week high
	FiftyTwoWeekLow            float64 // 52-week low
	FiftyDayMovingAverage      float64 // 50-day moving average
	TwoHundredDayMovingAverage float64 // 200-day moving average
	AvgVolume3Month            string  // Average volume (3 months)
	AvgVolume10Day             string  // Average volume (10 days)
	SharesOutstanding          float64 // Shares outstanding
	ImpliedSharesOutstanding   float64 // Implied shares outstanding
	Float                      float64 // Float
	PercentHeldByInsiders      float64 // % held by insiders
	PercentHeldByInstitutions  float64 // % held by institutions

	// Dividends & Splits
	ForwardAnnualDividendRate   float64 // Forward annual dividend rate
	ForwardAnnualDividendYield  float64 // Forward annual dividend yield
	TrailingAnnualDividendRate  float64 // Trailing annual dividend rate
	TrailingAnnualDividendYield float64 // Trailing annual dividend yield
	FiveYearAvgDividendYield    float64 // 5-year average dividend yield
	PayoutRatio                 float64 // Payout ratio
	ExDividendDate              string  // Ex-dividend date
	LastSplitFactor             string  // Last split factor
	LastSplitDate               string  // Last split date

	// Valuation Metrics
	ValuationMetrics ValuationMetrics
}

// ValuationMetrics holds valuation-related statistics
type ValuationMetrics struct {
	MarketCap              []string  // Market Cap for multiple periods
	EnterpriseValue        []string  // Enterprise Value for multiple periods
	TrailingPE             []float64 // Trailing P/E for multiple periods
	ForwardPE              []float64 // Forward P/E for multiple periods
	PEGRatio               []string  // PEG Ratio (5yr expected)
	PriceSales             []float64 // Price to Sales for multiple periods
	PriceBook              []float64 // Price to Book for multiple periods
	EnterpriseValueRevenue []float64 // Enterprise Value to Revenue for multiple periods
	EnterpriseValueEBITDA  []float64 // Enterprise Value to EBITDA for multiple periods
}

// Profile represents
type Profile struct {
	Name           string   // Stock name
	Sector         string   // Stock sector
	Description    string   // Stock name
	UpcomingEvents []string // Stock ticker symbol
	RecentEvents   []string // Stock sector
}

// IncomeStatement represents an income statement breakdown
type IncomeStatement struct {
	TotalRevenue                          []int64   // Total revenue in INR (thousands)
	CostOfRevenue                         []int64   // Cost of revenue
	GrossProfit                           []int64   // Gross profit
	OperatingExpense                      []int64   // Operating expenses
	OperatingIncome                       []int64   // Operating income
	NetNonOperatingInterestIncomeExpense  []int64   // Net non-operating interest income/expense
	OtherIncomeExpense                    []int64   // Other income/expense
	PretaxIncome                          []int64   // Pre-tax income
	TaxProvision                          []int64   // Tax provision
	EarningsFromEquityInterestNetOfTax    []int64   // Earnings from equity interest (net of tax)
	NetIncomeCommonStockholders           []int64   // Net income to common stockholders
	DilutedNIAvailableToComStockholders   []int64   // Diluted NI available to common stockholders
	BasicEPS                              []float64 // Basic EPS
	DilutedEPS                            []float64 // Diluted EPS
	BasicAverageShares                    []float64 // Basic average shares
	DilutedAverageShares                  []float64 // Diluted average shares
	TotalOperatingIncomeAsReported        []int64   // Total operating income as reported
	RentExpenseSupplemental               []int64   // Rent expense supplemental
	TotalExpenses                         []int64   // Total expenses
	NetIncomeFromContinuingDiscontinuedOp []int64   // Net income from continuing and discontinued operation
	NormalizedIncome                      []float64 // Normalized income
	InterestIncome                        []int64   // Interest income
	InterestExpense                       []int64   // Interest expense
	NetInterestIncome                     []int64   // Net interest income
	EBIT                                  []int64   // Earnings before interest and tax
	EBITDA                                []int64   // Earnings before interest, tax, depreciation, and amortization
	ReconciledCostOfRevenue               []int64   // Reconciled cost of revenue
	ReconciledDepreciation                []int64   // Reconciled depreciation
	NetIncomeFromContinuingOperation      []int64   // Net income from continuing operation
	TotalUnusualItemsExcludingGoodwill    []int64   // Total unusual items excluding goodwill
	TotalUnusualItems                     []int64   // Total unusual items
	NormalizedEBITDA                      []int64   // Normalized EBITDA
	TaxRateForCalcs                       []float64 // Tax rate for calculations
	TaxEffectOfUnusualItems               []float64 // Tax effect of unusual items
}

// BalanceSheet represents the entire balance sheet data
type BalanceSheet struct {
	Annual    BalanceSheetHistory // Annual data breakdown
	Quarterly BalanceSheetHistory // Quarterly data breakdown
}

// BalanceSheetHistory represents the historical data for a specific category
type BalanceSheetHistory struct {
	TotalAssets                 []int64   // Historical total assets for 4 periods
	TotalLiabilitiesNetMinority []int64   // Historical liabilities for 4 periods
	TotalEquityGrossMinority    []int64   // Historical equity for 4 periods
	TotalCapitalization         []int64   // Historical capitalization for 4 periods
	CommonStockEquity           []int64   // Historical common stock equity for 4 periods
	CapitalLeaseObligations     []int64   // Historical capital lease obligations for 4 periods
	NetTangibleAssets           []int64   // Historical net tangible assets for 4 periods
	WorkingCapital              []int64   // Historical working capital for 4 periods
	InvestedCapital             []int64   // Historical invested capital for 4 periods
	TangibleBookValue           []int64   // Historical tangible book value for 4 periods
	TotalDebt                   []int64   // Historical total debt for 4 periods
	NetDebt                     []int64   // Historical net debt for 4 periods
	ShareIssued                 []float64 // Historical shares issued for 4 periods
	OrdinarySharesNumber        []float64 // Historical ordinary shares for 4 periods
	TreasurySharesNumber        []float64 // Historical treasury shares for 4 periods
}

type Financial struct {
	BalanceSheet    BalanceSheet
	IncomeStatement IncomeStatement
}

// StockData represents stock-related information
type StockData struct {
	Profile        Profile
	News           []NewsItem        // List of news items
	Statistics     Statistics        // All statistical and financial information about the stock
	HistoricalData []Historical_data // histroical information
	Financial      Financial         // financial information
}

const (
	data_site = "https://finance.yahoo.com/quote/"
)

func executive_response_call(site string, selectors map[string]string) (error, map[string][]string) {
	// Create a new collector
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	// Map to store results for each selector
	data := make(map[string][]string)

	// Iterate over selectors and set up callbacks
	for key, selector := range selectors {
		// Register callback for each selector
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			data[key] = append(data[key], e.Text)
		})
	}

	// Handle errors
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	// Visit the provided site
	err := c.Visit(site)
	if err != nil {
		return err, nil
	}

	// Return extracted data
	return nil, data
}

func (s StockData) call_all_fetching_functions(s_code string) StockData {

	api_url := data_site + s_code

	fetch_profile_data(&s, api_url+"profile/")
	fetch_news_data(&s, api_url+"news")
	// fetch_staticts_data(&s,api_url+"key-statistics");
	// fetch_historicaldata_data(&s,api_url+"history");
	// fetch_financial_data(&s,api_url+"financials");
	return s
}

func fetch_profile_data(s *StockData, profile_site_api string) {
	// Define CSS selectors for required elements
	selectors := map[string]string{
		"Name":           "h1",
		"Sector":         "dd a",
		"Description":    "section.yf-1hj9jti p",
		"UpcomingEvents": "div.upcoming div.eventItem p",
		"RecentEvents":   "div.past div.eventItem p",
	}

	// Call the executive_response_call function
	err, data := executive_response_call(profile_site_api, selectors)
	if err != nil {
		fmt.Println("Error scraping data:", err)
		return
	}

	s.Profile.Name = data["Name"][1]
	s.Profile.Sector = data["Sector"][0]
	s.Profile.Description = data["Description"][0]
	s.Profile.UpcomingEvents = data["UpcomingEvents"]
	s.Profile.RecentEvents = data["RecentEvents"]
}

// Extract news data from the raw map returned by the scraper
func extract_news_data(data map[string][]string) ([]NewsItem, error) {
	// Create a slice to store news items
	var newsItems []NewsItem

	// Assuming the map contains "Title", "Summary", "Ticker", "Time", and "Image"
	for i := 0; i < len(data["Title"]); i++ {
		newsItem := NewsItem{
			Title:   data["Title"][i],
			Summary: data["Summary"][i],
			Ticker:  data["ticker"][i],
			Time:    data["Time"][i],
			Image:   data["Image"][i],
		}
		// Append the created news item to the slice
		newsItems = append(newsItems, newsItem)
	}

	return newsItems, nil
}

func fetch_news_data(s *StockData, news_site_api string) {
	// Define CSS selectors for required elements
	selectors := map[string]string{
		"Title":   "li.stream-item h3",
		"Summary": "li.stream-item p",
		"ticker":  "div.footer div.taxonomy-links a",
		"Time":    "div.footer div.publishing",
		"Image":   "li.stream-item a.thumb",
	}

	// Call the executive_response_call function to scrape the data
	err, data := executive_response_call(news_site_api, selectors)
	if err != nil {
		fmt.Println("Error scraping data:", err)
		return
	}

	// Extract news items from the raw data
	newsItems, err := extract_news_data(data)
	if err != nil {
		fmt.Println("Error extracting news:", err)
		return
	}

	// Assuming data.News is correctly populated, we now append it to s.News
	// Parse each item from the fetched data and append it to s.News
	for _, newsItemData := range newsItems {
		// Assuming newsItemData is a struct that has the necessary fields
		newsItem := NewsItem{
			Title:   newsItemData.Title,
			Summary: newsItemData.Summary,
			Ticker:  newsItemData.Ticker,
			Time:    newsItemData.Time,
			Image:   newsItemData.Image,
		}
		// Append the parsed news item to s.News
		s.News = append(s.News, newsItem)
	}

	// Optionally print the updated s.News to verify
	for _, newsItem := range s.News {
		fmt.Println(newsItem.Title)
	}
}

func fetch_staticts_data(s *StockData, staticts_site_api string) {

}

func fetch_historicaldata_data(s *StockData, historical_site_api string) {

}

func fetch_financial_data(s *StockData, financial_site_api string) {

}

func main() {
	// Create an instance of Data
	myData := StockData{}
	Data := myData.call_all_fetching_functions("VEDL.NS/")
	fmt.Println("originaly checking", Data)
}
