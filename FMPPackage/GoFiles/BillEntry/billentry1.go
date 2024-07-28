package billentry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BillStock(w http.ResponseWriter, r *http.Request) {

	fmt.Println("BEBillStock(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "GET" {

		var lSaveBillRespRec BillStockRespStruct

		lDetails, lFirstBill_No, err := GetBillStock()
		if err != nil {
			lSaveBillRespRec.ErrMsg = "BEBS01" + err.Error()
			lSaveBillRespRec.Status = "E"
		} else {
			if lDetails != nil {
				lSaveBillRespRec.BillStockArr = lDetails
				lSaveBillRespRec.Bill_No = lFirstBill_No + 1
				lSaveBillRespRec.Msg = "Bill Stock Received Successfully"
				lSaveBillRespRec.Status = "S"
			} else {
				lSaveBillRespRec.Msg = "Record not Found"
				lSaveBillRespRec.Status = "E"

			}
		}

		datas, err := json.Marshal(lSaveBillRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("BEBillStock(-)")

	}
}

func AddBill(w http.ResponseWriter, r *http.Request) {

	fmt.Println("BEAddBill(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "PUT" {

		var lBillDetailsRec BillStockStruct

		var lSaveBillRespRec AddMedRespStruct
		//Read and UnMarShall

		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			lSaveBillRespRec.ErrMsg = "BEAB01" + err.Error()
			lSaveBillRespRec.Status = "E"
		} else {

			err := json.Unmarshal(body, &lBillDetailsRec)
			fmt.Println(lBillDetailsRec)
			if err != nil {
				lSaveBillRespRec.ErrMsg = "BEAB02" + err.Error()
				lSaveBillRespRec.Status = "E"
			} else {

				//Fetch Data
				lDetails, err := AddBillMedicine(lBillDetailsRec)
				if err != nil {
					lSaveBillRespRec.ErrMsg = "BEAB01" + err.Error()
					lSaveBillRespRec.Status = "E"
				} else {
					if lDetails != nil {
						lSaveBillRespRec.MedArr = lDetails
						lSaveBillRespRec.Msg = "Add Med Received Successfully"
						lSaveBillRespRec.Status = "S"
					} else {
						lSaveBillRespRec.Msg = "Record not Found"
						lSaveBillRespRec.Status = "E"

					}
				}
			}
		}
		//Marshal
		datas, err := json.Marshal(lSaveBillRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("BEAddBill(-)")

	}
}

func SaveBill(w http.ResponseWriter, r *http.Request) {

	fmt.Println("BESaveBill(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	if r.Method == "PUT" {

		var lBillDetailsRec BillSaveReqStruct
		var lSaveBillRespRec BillSaveRespstruct
		//Read and UnMarShall

		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			lSaveBillRespRec.ErrMsg = "BESB01" + err.Error()
			lSaveBillRespRec.Status = "E"
		} else {

			err := json.Unmarshal(body, &lBillDetailsRec)
			fmt.Println(lBillDetailsRec)
			if err != nil {
				lSaveBillRespRec.ErrMsg = "BESB02" + err.Error()
				lSaveBillRespRec.Status = "E"
			} else {

				//Fetch Data
				lDetails, err := InsertBillDetails(lBillDetailsRec)
				if err != nil {
					lSaveBillRespRec.ErrMsg = "BESB01" + err.Error()
					lSaveBillRespRec.Status = "E"
				} else {
					if lDetails.Total_Price != 0 {
						BillPrice := lDetails

						lDetails2, err := InsertBillMaster(lBillDetailsRec, BillPrice)
						if err != nil {
							lSaveBillRespRec.ErrMsg = "BESB02" + err.Error()
							lSaveBillRespRec.Status = "E"
						} else {
							if lDetails2 != 0 {
								lSaveBillRespRec.Bill_No = lDetails2
								lSaveBillRespRec.Msg = "Bill Num Received"
								lSaveBillRespRec.Status = "S"
							} else {
								lSaveBillRespRec.Status = "E"
								lSaveBillRespRec.Msg = "Record Not Found"
							}
						}
					}

				}
			}
		}
		//Marshal
		datas, err := json.Marshal(lSaveBillRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data", err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("BESaveBill(-)")

	}
}
