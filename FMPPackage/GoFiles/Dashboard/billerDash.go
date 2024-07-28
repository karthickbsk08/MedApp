package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func BillerDash(w http.ResponseWriter, r *http.Request) {

	fmt.Println("DBBillerDash(+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "LOGINID,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {

		var lBillerDashrespRec BillerDashRespStruct

		//Reading and Unmarshal

		value := r.Header.Get("LOGINID")
		lLogin_id, err := strconv.Atoi(value)
		if err != nil {
			lBillerDashrespRec.ErrMsg = "DBBD01" + err.Error()
			lBillerDashrespRec.Status = "E"
		} else {
			lDetails, err := BillerTodaySales(lLogin_id)
			if err != nil {
				lBillerDashrespRec.ErrMsg = "DBBD02" + err.Error()
				lBillerDashrespRec.Status = "E"
			} else {

				lBillerDashrespRec.Today_Sales = lDetails
				lDetails1, err := BillerYesterdaySales(lLogin_id)
				if err != nil {
					lBillerDashrespRec.ErrMsg = "DBBD03" + err.Error()
					lBillerDashrespRec.Status = "E"
				} else {
					lBillerDashrespRec.Yesterday_Sales = lDetails1
					lBillerDashrespRec.Status = "S"
					lBillerDashrespRec.Msg = "Record Received successfull"

				}
			}
		}

		//UnMarshal
		datas, err := json.Marshal(lBillerDashrespRec)
		if err != nil {
			fmt.Fprintf(w, "Error Taking Data"+err.Error())
		} else {
			fmt.Fprintf(w, string(datas))
		}

		fmt.Println("DBBillerDash(-)")
	}
}
