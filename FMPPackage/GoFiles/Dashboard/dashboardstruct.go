package dashboard

//biller Dash
type BillerDashRespStruct struct {
	Today_Sales     float64 `json:"todaysales"`
	Yesterday_Sales float64 `json:"yesterdaysales"`
	ErrMsg          string  `json:"errmsg"`
	Status          string  `json:"status"`
	Msg             string  `json:"msg"`
}

type ManagerDashRespStruct struct {
	OAToday_sales   float64                 `json:"oatoday_Sales"`
	Inventory_value float64                 `json:"inventory_Value"`
	DailySalesArr   []DailySalesStruct      `json:"dailysalesarr"`
	MonthlySalesArr []MonthlySalesStruct    `json:"monthlysalesarr"`
	CurrentMonthArr []CurrentMonthPefStruct `json:"currentmontharr"`
	TodaySalesArr   []TodaySalesPefStruct   `json:"todaysalesarr"`
	Error           []ErrorStruct           `json:"error"`
	Status          string                  `json:"status"`
	Msg             string                  `json:"msg"`
}

// type DailyStruct struct {
// 	DailySalesArr []DailySalesStruct `json:"dailysalesarr"`
// 	Status        string             `json:"status"`
// 	Msg           string             `json:"msg"`
// }

// type common struct{
// 	ar interface{}
// 	Status
// 	masg

// }

type ErrorStruct struct {
	ErrorCode string `json:"errorcode"`
	ErrMsg    string `json:"errmsg"`
}

type DailySalesStruct struct {
	Total_amount float64 `json:"total_amount"`
	Day          string  `json:"day"`
}

type MonthlySalesStruct struct {
	Total_amount float64 `json:"total_amount"`
	Month        string  `json:"month"`
}

type CurrentMonthPefStruct struct {
	Total_amount float64 `json:"total_amount"`
	User_id      string  `json:"user_id"`
}

type TodaySalesPefStruct struct {
	Total_amount float64 `json:"total_amount"`
	User_id      string  `json:"user_id"`
}
