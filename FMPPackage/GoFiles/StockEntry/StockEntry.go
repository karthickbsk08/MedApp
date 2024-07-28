package stockentry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	common "medapp/FMPPackage/Common"
	"net/http"
)

type AddMedicineReqStruct struct {
	Medicine_Name string `json:"medicine_name"`
	Brand         string `json:"brand"`
	User_id       string `json:"user_id"`
}
type AddMedStockRespStruct struct {
	Last_Medicine common.StockStruct `json:"last_medicine"`
	ErrMsg        string             `json:"errmsg"`
	Status        string             `json:"status"`
	Msg           string             `json:"msg"`
}

func AddMedicine(w http.ResponseWriter, r *http.Request) {

	fmt.Println("SEAddMedicine(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "PUT" {

		var lMedicineRec AddMedicineReqStruct

		var lStocksRespRec AddMedStockRespStruct

		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			lStocksRespRec.ErrMsg = "SEAM01" + err.Error()
			lStocksRespRec.Status = "E"
		} else {

			err := json.Unmarshal(body, &lMedicineRec)
			fmt.Println(lMedicineRec)
			if err != nil {
				lStocksRespRec.ErrMsg = "SEAM02" + err.Error()
				lStocksRespRec.Status = "E"
			} else {
				lDetails, err := InsertNewMedicine(lMedicineRec)
				if err != nil {
					lStocksRespRec.ErrMsg = "SEAM03" + err.Error()
					lStocksRespRec.Status = "E"
				} else {
					if lDetails.Medicine_Name != "" {
						lStocksRespRec.Last_Medicine = lDetails
						lStocksRespRec.Msg = "Added Successfully"
						lStocksRespRec.Status = "S"
					} else {
						lStocksRespRec.Msg = "Already Exists"
						lStocksRespRec.Status = "E"

					}
				}
			}
		}
		datas, err := json.Marshal(lStocksRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("SEAddMedicine(-)")

	}
}

// type StockStruct struct {
// 	Medicine_Name string `json:"medicine_name"`
// 	Brand         string `json:"brand"`
// }

type StockRespStruct struct {
	StockArr []common.StockStruct `json:"stockarr"`
	ErrMsg   string               `json:"errmsg"`
	Status   string               `json:"status"`
	Msg      string               `json:"msg"`
}

func GetStocks(w http.ResponseWriter, r *http.Request) {

	fmt.Println("SEGetStocks(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "GET" {

		var lStocksRespRec StockRespStruct

		lDetails, err := FetchStocks()
		if err != nil {
			lStocksRespRec.ErrMsg = "SEAM03" + err.Error()
			lStocksRespRec.Status = "E"
		} else {
			if lDetails != nil {
				lStocksRespRec.StockArr = lDetails
				lStocksRespRec.Msg = "Stock Received Successfully"
				lStocksRespRec.Status = "S"
			} else {
				lStocksRespRec.Msg = "Record not Found"
				lStocksRespRec.Status = "E"

			}
		}

		datas, err := json.Marshal(lStocksRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("SEGetStocks(-)")

	}
}

type UpdateStockReqStruct struct {
	Medicine_Name string  `json:"medicine_name"`
	Brand         string  `json:"brand"`
	Quantity      int     `json:"quantity"`
	Unit_price    float64 `json:"unit_price"`
	User_id       string  `json:"user_id"`
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

	fmt.Println("SEUpdateStock(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "PUT" {

		var lUpdateMedRec UpdateStockReqStruct

		var lfinalrespRec common.CommonRespStruct

		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			lfinalrespRec.ErrMsg = "SEUS01" + err.Error()
			lfinalrespRec.Status = "E"
		} else {

			err := json.Unmarshal(body, &lUpdateMedRec)
			fmt.Println(lUpdateMedRec)
			if err != nil {
				lfinalrespRec.ErrMsg = "SEUS02" + err.Error()
				lfinalrespRec.Status = "E"
			} else {
				lDetails, err := StockUpdate(lUpdateMedRec)
				if err != nil {
					lfinalrespRec.ErrMsg = "SEUS03" + err.Error()
					lfinalrespRec.Status = "E"
				} else {
					if lDetails != 0 {
						lfinalrespRec.Msg = "UpDated Successfully"
						lfinalrespRec.Status = "S"
					} else {
						lFinalDetails, err := StockInsert(lUpdateMedRec)
						if err != nil {
							lfinalrespRec.ErrMsg = "SEUS03" + err.Error()
							lfinalrespRec.Status = "E"
						} else {

							if lFinalDetails != 0 {
								lfinalrespRec.Msg = "Updated Successfully"
								lfinalrespRec.Status = "S"
							} else {
								lfinalrespRec.Msg = "Stock Not Updated"
								lfinalrespRec.Status = "E"
							}

						}
					}
				}
			}
		}

		datas, err := json.Marshal(lfinalrespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("SEUpdateStock(-)")

	}
}
