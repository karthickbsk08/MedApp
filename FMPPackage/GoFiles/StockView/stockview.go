package stockview

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StockviewStruct struct {
	Medicine_Name string  `json:"medicine_name"`
	Brand         string  `json:"brand"`
	Quantity      int     `json:"quantity"`
	Unit_price    float64 `json:"unit_price"`
}
type StockviewRespStruct struct {
	StockviewArr []StockviewStruct `json:"stockviewarr"`
	ErrMsg       string            `json:"errmsg"`
	Status       string            `json:"status"`
	Msg          string            `json:"msg"`
}

func Stocksview(w http.ResponseWriter, r *http.Request) {

	fmt.Println("SVStocksview(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "GET" {

		var lStockiewRespRec StockviewRespStruct

		//Stocksview Fetch
		lDetails, err := FetchStockviewRecords()
		if err != nil {
			lStockiewRespRec.ErrMsg = "LHL01" + err.Error()
			lStockiewRespRec.Status = "E"

		} else {

			if lDetails[0].Medicine_Name != "" {
				lStockiewRespRec.StockviewArr = lDetails
				lStockiewRespRec.Status = "S"
				lStockiewRespRec.Msg = "Stock view record received"
			} else {
				lStockiewRespRec.Msg = "Record not Found"
				lStockiewRespRec.Status = "E"
			}
		}

		//UnMarshal
		datas, err := json.Marshal(lStockiewRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}
		fmt.Println("SVStocksview(-)")

	}
}
