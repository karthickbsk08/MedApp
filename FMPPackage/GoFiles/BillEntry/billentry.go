package billentry

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BillStockRespStruct1 struct {
	BillStockArr []BillStockStruct1 `json:"billstockarr"`
	ErrMsg       string             `json:"errmsg"`
	Status       string             `json:"status"`
	Msg          string             `json:"msg"`
}

type BillStockStruct1 struct {
	Medicine_Name string  `json:"medicine_name"`
	Quantity      int     `json:"quantity"`
	Brand         string  `json:"brand"`
	Unit_price    float64 `json:"unit_price"`
	Medicine_id   int     `json:"medicine_master_id"`
}

func BillStock1(w http.ResponseWriter, r *http.Request) {

	fmt.Println("BEBillStock(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "GET" {

		var lBillStockRespRec BillStockRespStruct1

		lDetails, err := GetBillStock1()
		if err != nil {
			lBillStockRespRec.ErrMsg = "BEBS01" + err.Error()
			lBillStockRespRec.Status = "E"
		} else {
			if lDetails != nil {
				lBillStockRespRec.BillStockArr = lDetails
				lBillStockRespRec.Msg = "Bill Stock Received Successfully"
				lBillStockRespRec.Status = "S"
			} else {
				lBillStockRespRec.Msg = "Record not Found"
				lBillStockRespRec.Status = "E"

			}
		}

		datas, err := json.Marshal(lBillStockRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("BEBillStock(-)")

	}
}
