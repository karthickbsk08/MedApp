package salesreport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SalesReportReqStruct struct {
	From_date string `json:"from_date"`
	To_date   string `json:"to_date"`
}

type SalesReportStruct struct {
	Bill_No       string  `json:"bill_no"`
	Bill_date     string  `json:"bill_date"`
	Medicine_Name string  `json:"medicine_name"`
	Quantity      int     `json:"quantity"`
	Amount        float64 `json:"amount"`
}

type SalesReportRespStruct struct {
	SalesArr []SalesReportStruct `json:"salesarr"`
	ErrMsg   string              `json:"errmsg"`
	Status   string              `json:"status"`
	Msg      string              `json:"msg"`
}

func GetSalesReport(w http.ResponseWriter, r *http.Request) {

	fmt.Println("SRGetSalesReport(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "PUT" {

		var lDateRangeRec SalesReportReqStruct

		var lSalesfinalRespRec SalesReportRespStruct

		//Read and Unmarshal
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			lSalesfinalRespRec.ErrMsg = "SRGS01" + err.Error()
			lSalesfinalRespRec.Status = "E"
		} else {
			err := json.Unmarshal(body, &lDateRangeRec)
			if err != nil {
				lSalesfinalRespRec.ErrMsg = "SRGS02" + err.Error()
				lSalesfinalRespRec.Status = "E"
			} else {
				var lDetails = make([]SalesReportStruct, 0)
				//SalesReport Fetch
				lDetails, err = FetchSalesReport(lDateRangeRec)
				fmt.Println(lDetails)
				if err != nil {
					lSalesfinalRespRec.ErrMsg = "SRGS03" + err.Error()
					lSalesfinalRespRec.Status = "E"

				}
				// else {

				// if lDetails != nil {
				// 	lSalesfinalRespRec.SalesArr = lDetails
				// 	lSalesfinalRespRec.Status = "S"
				// 	lSalesfinalRespRec.Msg = "Sales Report record received"
				// } else {
				// 	fmt.Println(lDetails)
				// 	lSalesfinalRespRec.Msg = "Record not Found"
				// 	lSalesfinalRespRec.Status = "E"
				// }
				// }

			}
		}
		//UnMarshal
		datas, err := json.Marshal(lSalesfinalRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}
		fmt.Println("SRGetSalesReport(-)")
	}

}
