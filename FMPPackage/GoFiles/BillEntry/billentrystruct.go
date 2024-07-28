package billentry

//Bill Stock Api
type BillStockRespStruct struct {
	BillStockArr []BillStockStruct `json:"billstockarr"`
	Bill_No      int               `json:"bill_no"`
	ErrMsg       string            `json:"errmsg"`
	Status       string            `json:"status"`
	Msg          string            `json:"msg"`
}

type BillStockStruct struct {
	Medicine_Name string `json:"medicine_name"`
	Quantity      int    `json:"quantity"`
}

//Add Medicine API

type AddMedRespStruct struct {
	MedArr []MedStruct `json:"medicinearr"`
	ErrMsg string      `json:"errmsg"`
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
}

type MedStruct struct {
	Medicine_id   int     `json:"medicine_master_id"`
	Medicine_Name string  `json:"medicine_name"`
	Brand         string  `json:"brand"`
	Quantity      int     `json:"quantity"`
	Amount        float64 `json:"amount"`
}

// type AddBillReqStruct struct {
// 	Medicine_Name string `json:"medicine_name"`
// 	Quantity      int    `json:"quantity"`
// }

//Save Bill API

type BillPriceStruct struct {
	Total_Price float64 `json:"bill_amount"`
	Gst         float64 `json:"bill_gst"`
	Net_Price   float64 `json:"net_price"`
}
type BillSaveReqStruct struct {
	User_id   string      `json:"user_id"`
	Bill_No   int         `json:"bill_no"`
	Bill_date string      `json:"bill_date"`
	Login_id  int         `json:"login_id"`
	BillsArr  []MedStruct `json:"billsarr"`
}

type BillSaveRespstruct struct {
	Bill_No int    `json:"bill_no"`
	ErrMsg  string `json:"errmsg"`
	Status  string `json:"status"`
	Msg     string `json:"msg"`
}
